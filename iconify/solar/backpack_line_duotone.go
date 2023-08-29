package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackpackLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 10.911v-.18a6 6 0 0 1 4.618-5.757l.176-.04l.167-.036a19 19 0 0 1 8.078 0l.167.037l.176.04A6 6 0 0 1 21 10.91v5.464a4.519 4.519 0 0 1-3.538 4.411c-3.598.8-7.326.8-10.923 0A4.519 4.519 0 0 1 3 16.376V10.91Z"/><path stroke-linecap="round" d="M17.5 15.5V17M15.959 4.5A3 3 0 0 0 13 2h-2a3 3 0 0 0-2.959 2.5" opacity=".5"/><path stroke-linecap="round" d="M3 14a22.16 22.16 0 0 0 18 0"/><path stroke-linecap="round" d="M10 13h4" opacity=".5"/></g>`),
		g.Group(children),
	)
}