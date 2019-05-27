# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import oneshot_pb2 as oneshot__pb2


class OneshotSuggestionStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetSuggestions = channel.unary_stream(
        '/api.v1.alpha3.OneshotSuggestion/GetSuggestions',
        request_serializer=oneshot__pb2.GetOneshotSuggestionRequest.SerializeToString,
        response_deserializer=oneshot__pb2.GetOneshotSuggestionReply.FromString,
        )


class OneshotSuggestionServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetSuggestions(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_OneshotSuggestionServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetSuggestions': grpc.unary_stream_rpc_method_handler(
          servicer.GetSuggestions,
          request_deserializer=oneshot__pb2.GetOneshotSuggestionRequest.FromString,
          response_serializer=oneshot__pb2.GetOneshotSuggestionReply.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'api.v1.alpha3.OneshotSuggestion', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))