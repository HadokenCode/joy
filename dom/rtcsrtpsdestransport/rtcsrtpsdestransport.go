package rtcsrtpsdestransport

import (
	"github.com/matthewmueller/joy/dom/rtcicetransport"
	"github.com/matthewmueller/joy/dom/rtcsrtpsdesparameters"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.EventTarget = (*RTCSrtpSdesTransport)(nil)

// New fn
func New(transport *rtcicetransport.RTCIceTransport, encryptparameters *rtcsrtpsdesparameters.RTCSrtpSdesParameters, decryptparameters *rtcsrtpsdesparameters.RTCSrtpSdesParameters) *RTCSrtpSdesTransport {
	js.Rewrite("RTCSrtpSdesTransport")
	return &RTCSrtpSdesTransport{}
}

// RTCSrtpSdesTransport struct
// js:"RTCSrtpSdesTransport,omit"
type RTCSrtpSdesTransport struct {
}

// GetLocalParameters fn
// js:"getLocalParameters"
func (*RTCSrtpSdesTransport) GetLocalParameters() (r []*rtcsrtpsdesparameters.RTCSrtpSdesParameters) {
	js.Rewrite("$_.getLocalParameters()")
	return r
}

// AddEventListener fn
// js:"addEventListener"
func (*RTCSrtpSdesTransport) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*RTCSrtpSdesTransport) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*RTCSrtpSdesTransport) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Onerror prop
// js:"onerror"
func (*RTCSrtpSdesTransport) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*RTCSrtpSdesTransport) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// Transport prop
// js:"transport"
func (*RTCSrtpSdesTransport) Transport() (transport *rtcicetransport.RTCIceTransport) {
	js.Rewrite("$_.transport")
	return transport
}
