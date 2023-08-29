package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phplist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 384H704v256l-64 64l-64-64V384H448v256l-64-64l-64 64V384H192q-80 0-136-56T0 192.5t56-136T192 0h256v256h128V0h256q80 0 136 56t56 136t-56.5 136T832 384zM320 128H192q-27 0-45.5 19T128 192.5t18.5 45T192 256h128V128zm512 0H704v128h128q26 0 45-18.5t19-45t-19-45.5t-45-19z"/>`),
		g.Group(children),
	)
}