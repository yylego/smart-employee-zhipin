package main

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/yylego/done"
	"github.com/yylego/kratos-vue3/vue3kratos"
	"github.com/yylego/must"
	"github.com/yylego/osexec"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/rese"
	"github.com/yylego/runpath"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

func main() {
	zaplog.SUG.Infoln("=== Vue3 Client Code Gen Workflow Start ===")

	frontendRoot := runpath.PARENT.UpTo(1, "zhipin-vue3")
	kratosRoot := runpath.PARENT.UpTo(1, "zhipin-kratos")

	zaplog.LOG.Debug("paths", zap.String("backend", kratosRoot), zap.String("frontend", frontendRoot))
	runGenerate(kratosRoot, filepath.Join(frontendRoot, "src/rpc/zhipin"))

	zaplog.SUG.Infoln("=== WORKFLOW FINISHED SUCCESS! ===")
}

func runGenerate(kratosRoot string, clientCodeDest string) {
	zaplog.SUG.Infoln("Backend project:", kratosRoot)
	osmustexist.ROOT(kratosRoot)

	makefilePath := filepath.Join(kratosRoot, "Makefile")
	osmustexist.FILE(makefilePath)

	makefileData := rese.A1(os.ReadFile(makefilePath))
	must.True(bytes.Contains(makefileData, []byte("web_api_grpc_ts:")))
	must.True(bytes.Contains(makefileData, []byte("web_api_grpc_to_http:")))
	must.True(bytes.Contains(makefileData, []byte("web_api_cleanup:")))
	zaplog.SUG.Infoln("Makefile targets verified")

	grpcTsOutput := filepath.Join(kratosRoot, "bin", "web_api_grpc_ts.out")
	zaplog.SUG.Infoln("Generating TypeScript gRPC clients...")

	if osmustexist.IsRootExist(filepath.Join(kratosRoot, "bin")) {
		zaplog.SUG.Infoln("   Cleaning previous output...")
		done.Done(os.RemoveAll(grpcTsOutput))
	}

	rese.A1(osexec.ExecInPath(kratosRoot, "make", "web_api_grpc_ts"))
	osmustexist.ROOT(grpcTsOutput)
	zaplog.SUG.Infoln("TypeScript gRPC clients generated")

	zaplog.SUG.Infoln("Converting gRPC clients to HTTP clients...")
	rese.A1(osexec.ExecInPath(kratosRoot, "make", "web_api_grpc_to_http"))
	zaplog.SUG.Infoln("Conversion completed")

	// create dest dir if not exists
	must.Done(os.MkdirAll(clientCodeDest, 0755))

	zaplog.SUG.Infoln("Syncing converted files...")
	zaplog.SUG.Infoln("   From:", grpcTsOutput)
	zaplog.SUG.Infoln("   To:  ", clientCodeDest)
	vue3kratos.CloneFilesToDestRoot(grpcTsOutput, clientCodeDest)
	zaplog.SUG.Infoln("File sync completed")

	zaplog.SUG.Infoln("Cleaning up temp files...")
	rese.A1(osexec.ExecInPath(kratosRoot, "make", "web_api_cleanup"))
	zaplog.SUG.Infoln("Cleanup completed")
}
