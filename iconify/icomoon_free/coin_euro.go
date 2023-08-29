package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoinEuro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.5 1a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15zm0 13.5a6 6 0 1 1 0-12a6 6 0 0 1 0 12z"/><path fill="currentColor" d="M10.482 10.068a.5.5 0 0 0-.684.181A1.51 1.51 0 0 1 8.5 11h-2a1.502 1.502 0 0 1-1.414-1H7.5a.5.5 0 0 0 0-1H5V8h2.5a.5.5 0 0 0 0-1H5.086c.206-.582.762-1 1.414-1h2a1.51 1.51 0 0 1 1.298.751a.5.5 0 1 0 .865-.503a2.513 2.513 0 0 0-2.162-1.249h-2c-1.207 0-2.217.86-2.45 2h-.55a.5.5 0 0 0 0 1h.5v1h-.5a.5.5 0 0 0 0 1h.55c.232 1.14 1.242 2 2.45 2h2a2.51 2.51 0 0 0 2.162-1.249a.5.5 0 0 0-.181-.684z"/>`),
		g.Group(children),
	)
}