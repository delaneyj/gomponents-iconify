package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClearFormattingEraser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m704 1434l422 422l-65 64H731l-211-211q-19-19-19-45t19-45l184-185zm640-640l422 422l-550 550l-422-422l550-550z"/>`),
		g.Group(children),
	)
}