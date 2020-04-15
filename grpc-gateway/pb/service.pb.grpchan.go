// Code generated by protoc-gen-grpchan. DO NOT EDIT.
// source: pb/service.proto

package pb

import "github.com/fullstorydev/grpchan"
import "golang.org/x/net/context"
import "google.golang.org/grpc"

func RegisterHandlerGrpcGateway(reg grpchan.ServiceRegistry, srv GrpcGatewayServer) {
	reg.RegisterService(&_GrpcGateway_serviceDesc, srv)
}

type grpcGatewayChannelClient struct {
	ch grpchan.Channel
}

func NewGrpcGatewayChannelClient(ch grpchan.Channel) GrpcGatewayClient {
	return &grpcGatewayChannelClient{ch: ch}
}

func (c *grpcGatewayChannelClient) GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (GrpcGateway_GetDevicesClient, error) {
	stream, err := c.ch.NewStream(ctx, &_GrpcGateway_serviceDesc.Streams[0], "/ocf.cloud.grpcgateway.pb.GrpcGateway/GetDevices", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcGatewayGetDevicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

func (c *grpcGatewayChannelClient) GetResourceLinks(ctx context.Context, in *GetResourceLinksRequest, opts ...grpc.CallOption) (GrpcGateway_GetResourceLinksClient, error) {
	stream, err := c.ch.NewStream(ctx, &_GrpcGateway_serviceDesc.Streams[1], "/ocf.cloud.grpcgateway.pb.GrpcGateway/GetResourceLinks", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcGatewayGetResourceLinksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

func (c *grpcGatewayChannelClient) RetrieveResourceFromDevice(ctx context.Context, in *RetrieveResourceFromDeviceRequest, opts ...grpc.CallOption) (*RetrieveResourceFromDeviceResponse, error) {
	out := new(RetrieveResourceFromDeviceResponse)
	err := c.ch.Invoke(ctx, "/ocf.cloud.grpcgateway.pb.GrpcGateway/RetrieveResourceFromDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcGatewayChannelClient) RetrieveResourcesValues(ctx context.Context, in *RetrieveResourcesValuesRequest, opts ...grpc.CallOption) (GrpcGateway_RetrieveResourcesValuesClient, error) {
	stream, err := c.ch.NewStream(ctx, &_GrpcGateway_serviceDesc.Streams[2], "/ocf.cloud.grpcgateway.pb.GrpcGateway/RetrieveResourcesValues", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcGatewayRetrieveResourcesValuesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

func (c *grpcGatewayChannelClient) UpdateResourcesValues(ctx context.Context, in *UpdateResourceValuesRequest, opts ...grpc.CallOption) (*UpdateResourceValuesResponse, error) {
	out := new(UpdateResourceValuesResponse)
	err := c.ch.Invoke(ctx, "/ocf.cloud.grpcgateway.pb.GrpcGateway/UpdateResourcesValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcGatewayChannelClient) SubscribeForEvents(ctx context.Context, opts ...grpc.CallOption) (GrpcGateway_SubscribeForEventsClient, error) {
	stream, err := c.ch.NewStream(ctx, &_GrpcGateway_serviceDesc.Streams[3], "/ocf.cloud.grpcgateway.pb.GrpcGateway/SubscribeForEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcGatewaySubscribeForEventsClient{stream}
	return x, nil
}