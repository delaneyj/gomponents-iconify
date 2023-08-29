package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhitePaper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m22 27.18l-2.59-2.59L18 26l4 4l8-8l-1.41-1.41L22 27.18zM9 17h7v2H9zm0-5h12v2H9zm0-5h12v2H9z"/><path fill="currentColor" d="M16 30H6c-1.103 0-2-.897-2-2V4c0-1.103.897-2 2-2h18c1.103 0 2 .897 2 2v15h-2V4H6v24h10v2Z"/>`),
		g.Group(children),
	)
}