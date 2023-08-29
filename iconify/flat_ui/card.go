package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Card(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<circle cx="50" cy="50" r="50" fill="#F2F2F2"/><clipPath id="flatUiCard0"><circle cx="50" cy="50" r="50"/></clipPath><g fill-rule="evenodd" clip-path="url(#flatUiCard0)" clip-rule="evenodd"><path fill="#E74C3C" d="M-19 18h90a5 5 0 0 1 5 5v54a5 5 0 0 1-5 5h-90a5 5 0 0 1-5-5V23a5 5 0 0 1 5-5z"/><path fill="#2C3E50" d="M-24 24H76v13H-24V24z"/><path fill="#ECF0F1" d="M-16 43h62v14h-62zm71 0h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H55a3 3 0 0 1-3-3v-8a3 3 0 0 1 3-3z"/><path fill="#fff" d="M70.948 43.736L53.052 56.264A2.97 2.97 0 0 0 55 57h14a3 3 0 0 0 3-3v-8a2.98 2.98 0 0 0-1.052-2.264z"/><path fill="#C0392B" d="M31.5 72h38a2.5 2.5 0 1 1 0 5h-38a2.5 2.5 0 1 1 0-5z"/></g>`),
		g.Group(children),
	)
}