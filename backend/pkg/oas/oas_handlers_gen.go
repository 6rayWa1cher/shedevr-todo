// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// handleCreateTaskRequest handles createTask operation.
//
// Creates a new task.
//
// POST /tasks
func (s *Server) handleCreateTaskRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createTask"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/tasks"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreateTask",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "CreateTask",
			ID:   "createTask",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityRemoteUserAuth(ctx, "CreateTask", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "RemoteUserAuth",
					Err:              err,
				}
				if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
					recordError("Security:RemoteUserAuth", err)
				}
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
				recordError("Security", err)
			}
			return
		}
	}
	request, close, err := s.decodeCreateTaskRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *Task
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "CreateTask",
			OperationSummary: "Create a new task",
			OperationID:      "createTask",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *NewTask
			Params   = struct{}
			Response = *Task
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.CreateTask(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.CreateTask(ctx, request)
	}
	if err != nil {
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w, span); err != nil {
				recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w, span); err != nil {
			recordError("Internal", err)
		}
		return
	}

	if err := encodeCreateTaskResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleDeleteTaskRequest handles deleteTask operation.
//
// Deletes a specific task by its ID.
//
// DELETE /tasks/{id}
func (s *Server) handleDeleteTaskRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteTask"),
		semconv.HTTPMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/tasks/{id}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeleteTask",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "DeleteTask",
			ID:   "deleteTask",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityRemoteUserAuth(ctx, "DeleteTask", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "RemoteUserAuth",
					Err:              err,
				}
				if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
					recordError("Security:RemoteUserAuth", err)
				}
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
				recordError("Security", err)
			}
			return
		}
	}
	params, err := decodeDeleteTaskParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response DeleteTaskRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "DeleteTask",
			OperationSummary: "Delete a task",
			OperationID:      "deleteTask",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = DeleteTaskParams
			Response = DeleteTaskRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeleteTaskParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.DeleteTask(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.DeleteTask(ctx, params)
	}
	if err != nil {
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w, span); err != nil {
				recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w, span); err != nil {
			recordError("Internal", err)
		}
		return
	}

	if err := encodeDeleteTaskResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetTaskByIdRequest handles getTaskById operation.
//
// Retrieves a specific task by its ID.
//
// GET /tasks/{id}
func (s *Server) handleGetTaskByIdRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getTaskById"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/tasks/{id}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "GetTaskById",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "GetTaskById",
			ID:   "getTaskById",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityRemoteUserAuth(ctx, "GetTaskById", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "RemoteUserAuth",
					Err:              err,
				}
				if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
					recordError("Security:RemoteUserAuth", err)
				}
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
				recordError("Security", err)
			}
			return
		}
	}
	params, err := decodeGetTaskByIdParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response GetTaskByIdRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "GetTaskById",
			OperationSummary: "Get a task by ID",
			OperationID:      "getTaskById",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetTaskByIdParams
			Response = GetTaskByIdRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetTaskByIdParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetTaskById(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetTaskById(ctx, params)
	}
	if err != nil {
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w, span); err != nil {
				recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w, span); err != nil {
			recordError("Internal", err)
		}
		return
	}

	if err := encodeGetTaskByIdResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetTasksRequest handles getTasks operation.
//
// Retrieves a list of all tasks.
//
// GET /tasks
func (s *Server) handleGetTasksRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getTasks"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/tasks"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "GetTasks",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "GetTasks",
			ID:   "getTasks",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityRemoteUserAuth(ctx, "GetTasks", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "RemoteUserAuth",
					Err:              err,
				}
				if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
					recordError("Security:RemoteUserAuth", err)
				}
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
				recordError("Security", err)
			}
			return
		}
	}

	var response []Task
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "GetTasks",
			OperationSummary: "Get all tasks",
			OperationID:      "getTasks",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = []Task
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetTasks(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetTasks(ctx)
	}
	if err != nil {
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w, span); err != nil {
				recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w, span); err != nil {
			recordError("Internal", err)
		}
		return
	}

	if err := encodeGetTasksResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUpdateTaskRequest handles updateTask operation.
//
// Updates a specific task by its ID.
//
// PUT /tasks/{id}
func (s *Server) handleUpdateTaskRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateTask"),
		semconv.HTTPMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/tasks/{id}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "UpdateTask",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "UpdateTask",
			ID:   "updateTask",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityRemoteUserAuth(ctx, "UpdateTask", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "RemoteUserAuth",
					Err:              err,
				}
				if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
					recordError("Security:RemoteUserAuth", err)
				}
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
				recordError("Security", err)
			}
			return
		}
	}
	params, err := decodeUpdateTaskParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	request, close, err := s.decodeUpdateTaskRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response UpdateTaskRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "UpdateTask",
			OperationSummary: "Update a task",
			OperationID:      "updateTask",
			Body:             request,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *UpdateTask
			Params   = UpdateTaskParams
			Response = UpdateTaskRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUpdateTaskParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.UpdateTask(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.UpdateTask(ctx, request, params)
	}
	if err != nil {
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w, span); err != nil {
				recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w, span); err != nil {
			recordError("Internal", err)
		}
		return
	}

	if err := encodeUpdateTaskResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}
