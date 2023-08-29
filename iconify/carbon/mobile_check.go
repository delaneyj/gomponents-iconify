package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m20 27.18l-2.59-2.59L16 26l4 4l8-8l-1.41-1.41L20 27.18z"/><path fill="currentColor" d="M10 28V10h12v9h2V6a2.002 2.002 0 0 0-2-2H10a2.002 2.002 0 0 0-2 2v22a2.002 2.002 0 0 0 2 2h4v-2Zm0-22h12v2H10Z"/>`),
		g.Group(children),
	)
}