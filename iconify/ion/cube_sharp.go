package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M48 170v196.92L240 480V284L48 170zm224 310l192-113.08V170L272 284Zm176-122.36ZM448 144L256 32L64 144l192 112l192-112z"/>`),
		g.Group(children),
	)
}