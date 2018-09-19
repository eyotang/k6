/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2018 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/eyotang/k6/core"
	"github.com/eyotang/k6/core/local"
	"github.com/eyotang/k6/js"
	"github.com/eyotang/k6/lib"
	"github.com/eyotang/k6/lib/types"
	"github.com/manyminds/api2go/jsonapi"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	null "gopkg.in/guregu/null.v3"
)

func TestSetupData(t *testing.T) {
	t.Parallel()
	runner, err := js.New(
		&lib.SourceData{Filename: "/script.js", Data: []byte(`
			export function setup() {
				return {"v": 1};
			}

			export default function(data) {
				if (!data || data.v != 2) {
					throw new Error("incorrect data: " + JSON.stringify(data));
				}
			};

			export function teardown(data) {
				if (!data || data.v != 2) {
					throw new Error("incorrect teardown data: " + JSON.stringify(data));
				}
			}

		`)},
		afero.NewMemMapFs(),
		lib.RuntimeOptions{},
	)
	require.NoError(t, err)
	runner.SetOptions(lib.Options{
		Paused:          null.BoolFrom(true),
		VUs:             null.IntFrom(2),
		VUsMax:          null.IntFrom(2),
		Iterations:      null.IntFrom(3),
		SetupTimeout:    types.NullDurationFrom(1 * time.Second),
		TeardownTimeout: types.NullDurationFrom(1 * time.Second),
	})
	executor := local.New(runner)
	executor.SetRunSetup(false)
	engine, err := core.NewEngine(executor, runner.GetOptions())
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	errC := make(chan error)
	go func() { errC <- engine.Run(ctx) }()

	handler := NewHandler()

	checkSetup := func(method, body, expResult string) {
		rw := httptest.NewRecorder()
		handler.ServeHTTP(rw, newRequestWithEngine(engine, method, "/v1/setup", bytes.NewBufferString(body)))
		res := rw.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var doc jsonapi.Document
		assert.NoError(t, json.Unmarshal(rw.Body.Bytes(), &doc))
		if !assert.NotNil(t, doc.Data.DataObject) {
			return
		}
		assert.Equal(t, "setupData", doc.Data.DataObject.Type)
		assert.JSONEq(t, expResult, string(doc.Data.DataObject.Attributes))
	}

	checkSetup("GET", "", `{"data": null}`)
	checkSetup("POST", "", `{"data": {"v":1}}`)
	checkSetup("GET", "", `{"data": {"v":1}}`)
	checkSetup("PUT", `{"v":2, "test":"mest"}`, `{"data": {"v":2, "test":"mest"}}`)
	checkSetup("GET", "", `{"data": {"v":2, "test":"mest"}}`)

	engine.Executor.SetPaused(false)

	select {
	case <-time.After(10 * time.Second):
		cancel()
		t.Fatal("Test timed out")
	case err := <-errC:
		cancel()
		require.NoError(t, err)
	}
}
