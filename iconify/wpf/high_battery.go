package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighBattery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M5 5C3.3 5 2 6.3 2 8v1H1c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1h1v1c0 1.7 1.3 3 3 3h18c1.7 0 3-1.3 3-3V8c0-1.7-1.3-3-3-3H5zm0 2h1v12H5c-.6 0-1-.4-1-1v-3H2v-4h2V8c0-.6.4-1 1-1z"/>`),
		g.Group(children),
	)
}