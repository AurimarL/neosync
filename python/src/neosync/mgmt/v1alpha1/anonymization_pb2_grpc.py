# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from mgmt.v1alpha1 import anonymization_pb2 as mgmt_dot_v1alpha1_dot_anonymization__pb2


class AnonymizationServiceStub(object):
    """Service that transactionally anonymizes data, regardless of the connection type.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.AnonymizeMany = channel.unary_unary(
                '/mgmt.v1alpha1.AnonymizationService/AnonymizeMany',
                request_serializer=mgmt_dot_v1alpha1_dot_anonymization__pb2.AnonymizeManyRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_anonymization__pb2.AnonymizeManyResponse.FromString,
                _registered_method=True)
        self.AnonymizeSingle = channel.unary_unary(
                '/mgmt.v1alpha1.AnonymizationService/AnonymizeSingle',
                request_serializer=mgmt_dot_v1alpha1_dot_anonymization__pb2.AnonymizeSingleRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_anonymization__pb2.AnonymizeSingleResponse.FromString,
                _registered_method=True)


class AnonymizationServiceServicer(object):
    """Service that transactionally anonymizes data, regardless of the connection type.
    """

    def AnonymizeMany(self, request, context):
        """Anonymizes many JSON strings by applying specified transformation mappings. This is the bulk version of the `AnonymizeSingle` method.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AnonymizeSingle(self, request, context):
        """Anonymizes a single JSON strings by applying specified transformation mappings.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AnonymizationServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'AnonymizeMany': grpc.unary_unary_rpc_method_handler(
                    servicer.AnonymizeMany,
                    request_deserializer=mgmt_dot_v1alpha1_dot_anonymization__pb2.AnonymizeManyRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_anonymization__pb2.AnonymizeManyResponse.SerializeToString,
            ),
            'AnonymizeSingle': grpc.unary_unary_rpc_method_handler(
                    servicer.AnonymizeSingle,
                    request_deserializer=mgmt_dot_v1alpha1_dot_anonymization__pb2.AnonymizeSingleRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_anonymization__pb2.AnonymizeSingleResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'mgmt.v1alpha1.AnonymizationService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('mgmt.v1alpha1.AnonymizationService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class AnonymizationService(object):
    """Service that transactionally anonymizes data, regardless of the connection type.
    """

    @staticmethod
    def AnonymizeMany(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/mgmt.v1alpha1.AnonymizationService/AnonymizeMany',
            mgmt_dot_v1alpha1_dot_anonymization__pb2.AnonymizeManyRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_anonymization__pb2.AnonymizeManyResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def AnonymizeSingle(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/mgmt.v1alpha1.AnonymizationService/AnonymizeSingle',
            mgmt_dot_v1alpha1_dot_anonymization__pb2.AnonymizeSingleRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_anonymization__pb2.AnonymizeSingleResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
