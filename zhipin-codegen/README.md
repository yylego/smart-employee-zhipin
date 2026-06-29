# zhipin-codegen

Generates the Vue3 frontend TypeScript client from the backend proto.

One command runs the whole workflow — clean old output, buf-generate the gRPC TypeScript client, convert it into an HTTP client, sync into `zhipin-vue3/src/rpc/zhipin`, then clean up:

```bash
make gen
```

Run it when a proto changes.
