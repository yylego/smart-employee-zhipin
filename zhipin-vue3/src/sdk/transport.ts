import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'

export const zhipinTransport = new GrpcWebFetchTransport({
    baseUrl: '/zhipin-api',
})
