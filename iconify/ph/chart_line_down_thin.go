package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLineDownThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 208a4 4 0 0 1-4 4H32a4 4 0 0 1-4-4V48a4 4 0 0 1 8 0v14.34l60 60l29.17-29.17a4 4 0 0 1 5.66 0L196 158.34V128a4 4 0 0 1 8 0v40a4.13 4.13 0 0 1-.08.78a3.37 3.37 0 0 1-.1.34a2.8 2.8 0 0 1-.13.41a2.87 2.87 0 0 1-.2.39c-.05.1-.1.2-.16.3a4.19 4.19 0 0 1-1.11 1.11l-.31.16a3.48 3.48 0 0 1-.38.2a2.8 2.8 0 0 1-.41.13a3.37 3.37 0 0 1-.34.1a4.13 4.13 0 0 1-.78.08h-40a4 4 0 0 1 0-8h30.34L128 101.66l-29.17 29.17a4 4 0 0 1-5.66 0L36 73.66V204h188a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}