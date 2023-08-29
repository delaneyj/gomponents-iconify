package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaterialThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M9 15V9.564c0-.892.087-1.215.25-1.54c.164-.327.404-.583.71-.757c.305-.174.608-.267 1.444-.267h25.192c.836 0 1.14.093 1.445.267c.305.174.545.43.709.756c.163.326.25.65.25 1.54V15"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M5 15h38v26H5V15Z"/><path fill="currentColor" fill-rule="evenodd" d="M13 26a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5.57 40.39L15 30l5 4l6-7l16.394 13.39"/></g>`),
		g.Group(children),
	)
}