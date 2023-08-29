package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5.219 6.688L3.78 8.094L16 20.313l12.219-12.22l-1.438-1.405L16 17.468zm0 7L3.78 15.094L16 27.313l12.219-12.22l-1.438-1.405L16 24.468z"/>`),
		g.Group(children),
	)
}