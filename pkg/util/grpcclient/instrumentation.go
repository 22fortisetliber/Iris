package grpcclient

import (
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	// "github.com/weaveworks/common/middleware"
	"google.golang.org/grpc"

	"github.com/22fortisetliber/iris/pkg/util/grpcutil"
	irismiddleware "github.com/22fortisetliber/iris/pkg/util/middleware"
)

func Instrument(requestDuration *prometheus.HistogramVec) ([]grpc.UnaryClientInterceptor, []grpc.StreamClientInterceptor) {
	return []grpc.UnaryClientInterceptor{
			grpcutil.HTTPHeaderPropagationClientInterceptor,
			otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer()),
			// middleware.ClientUserHeaderInterceptor,
			irismiddleware.PrometheusGRPCUnaryInstrumentation(requestDuration),
		}, []grpc.StreamClientInterceptor{
			grpcutil.HTTPHeaderPropagationStreamClientInterceptor,
			otgrpc.OpenTracingStreamClientInterceptor(opentracing.GlobalTracer()),
			// middleware.StreamClientUserHeaderInterceptor,
			irismiddleware.PrometheusGRPCStreamInstrumentation(requestDuration),
		}
}
