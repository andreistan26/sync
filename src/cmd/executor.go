package cmd

import (
	"fmt"

	"github.com/andreistan26/sync/src/file_level"
	"github.com/andreistan26/sync/src/options"
	transport "github.com/andreistan26/sync/src/transfer_level"
)

func Execute(opts *options.Options) error {
	switch opts.ExType {
	case options.LOCAL_EX:
		return ExecuteHostExchange(opts)
	case options.TCP_EX:
		return ExecuteTCPExchange(opts)
	default:
		return nil
	}
}

func ExecuteHostExchange(opts *options.Options) error {
	sf := file_level.CreateSourceFile(opts.Source.Filepath)
	rf := file_level.CreateRemoteFile(opts.Dest.Filepath)
	ex, err := file_level.CreateRsyncExchange(&sf, rf.ChunkList)

	file_level.CheckErr(err)

	resp := ex.Search()

	rf.WriteSyncedFile(&resp, opts.Dest.Filepath, true)

	return nil
}

func ExecuteTCPExchange(opts *options.Options) error {
	opts.Dest.Address += fmt.Sprintf(":%d", opts.Port)
	return transport.SendFile(opts)
}

func ExecuteStartServer(opts *options.ServerOptions) error {
	serv, err := transport.StartServer(options.DEFAULT_PORT)
	if err != nil {
		panic(err)
	}
	return serv.Run()
}
