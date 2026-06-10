package main

import (
	"flag"
	"net/http"
	"sync"
	"text/template"

	"github.com/gorilla/websocket"
	"github.com/pion/logging"
	"github.com/pion/webrtc/v4"
)

var (
	addr     = flag.String("addr", ":8080", "http service address")
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	indexTemplate = &template.Template{}

	// lock for peerConnections and trackLocals
	listLock        sync.RWMutex
	peerConnections []peerConnectionState
	trackLocals     map[string]*webrtc.TrackLocalStaticRTP

	log = logging.NewDefaultLoggerFactory().NewLogger("sfu-ws")
)
