package push

// import (
// 	"context"
// 	"encoding/json"
// 	"io"
// 	"net/http"
//
// 	"github.com/go-kit/log/level"
// 	"github.com/weaveworks/common/httpgrpc"
// 	"github.com/weaveworks/common/middleware"
//
// 	// "github.com/22fortisetliber/iris/pkg/irispb"
// 	"github.com/22fortisetliber/iris/pkg/util/log"
// )
//
// // Func defines the type of the push. It is similar to http.HandlerFunc.
// type Func func(context.Context, *irispb.WriteRequest) (*irispb.WriteResponse, error)
//
// // Handler is a http.Handler which accepts WriteRequests.
// func Handler(maxRecvMsgSize int, sourceIPs *middleware.SourceIPExtractor, push Func) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		ctx := r.Context()
// 		logger := log.WithContext(ctx, log.Logger)
//
// 		// var req irispb.WriteRequest
// 		var ag irispb.AlertGroup
// 		b, err := io.ReadAll(r.Body)
//
// 		if err != nil {
// 			level.Error(logger).Log("msg", "reading request body unsuccessfully", "err", err.Error())
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
//
// 		if err := json.Unmarshal(b, &ag); err != nil {
// 			level.Error(logger).Log("msg", "fail to unmarshall request body", "err", err.Error())
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
//
// 		req, err := irispb.ToProto(&ag)
// 		if err != nil {
// 			level.Error(logger).Log("msg", "fail to parse to proto", "err", err.Error())
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 		level.Debug(logger).Log("msg", "push request", "req", *req)
// 		if _, err := push(ctx, req); err != nil {
// 			resp, ok := httpgrpc.HTTPResponseFromError(err)
// 			if !ok {
// 				http.Error(w, err.Error(), http.StatusInternalServerError)
// 				return
// 			}
// 			if resp.GetCode()/100 == 5 {
// 				level.Error(logger).Log("msg", "push error", "err", err)
// 			} else if resp.GetCode() != http.StatusAccepted && resp.GetCode() != http.StatusTooManyRequests {
// 				level.Warn(logger).Log("msg", "push refused", "err", err)
// 			}
// 			http.Error(w, string(resp.Body), int(resp.Code))
// 		}
// 	})
// }
