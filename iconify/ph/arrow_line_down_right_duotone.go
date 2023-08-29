package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineDownRightDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M192 104v96H96l96-96z" opacity=".2"/><path d="M224 40a8 8 0 0 1-8 8H40a8 8 0 0 1 0-16h176a8 8 0 0 1 8 8Zm-24 64v96a8 8 0 0 1-8 8H96a8 8 0 0 1-5.66-13.66L132.69 152L74.34 93.66a8 8 0 0 1 11.32-11.32L144 140.69l42.34-42.35A8 8 0 0 1 200 104Zm-16 19.31l-34.34 34.35L115.31 192H184Z"/></g>`),
		g.Group(children),
	)
}