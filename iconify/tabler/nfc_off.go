package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NfcOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 20a3 3 0 0 1-3-3V8m5-4a3 3 0 0 1 3 3v5m0 4v2l-5-5"/><path d="M8 4h9a3 3 0 0 1 3 3v9m-.873 3.116A2.99 2.99 0 0 1 17 20H7a3 3 0 0 1-3-3V7c0-.83.337-1.582.882-2.125M3 3l18 18"/></g>`),
		g.Group(children),
	)
}