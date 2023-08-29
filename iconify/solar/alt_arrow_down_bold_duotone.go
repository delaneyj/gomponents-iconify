package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AltArrowDownBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m8.303 12.404l3.327 3.431c.213.22.527.22.74 0l6.43-6.63C19.201 8.79 18.958 8 18.43 8h-5.723l-4.404 4.404Z"/><path d="M11.293 8H5.57c-.528 0-.771.79-.37 1.205l2.406 2.481L11.293 8Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}