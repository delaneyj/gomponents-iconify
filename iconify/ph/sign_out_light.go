package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignOutLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M110 216a6 6 0 0 1-6 6H48a14 14 0 0 1-14-14V48a14 14 0 0 1 14-14h56a6 6 0 0 1 0 12H48a2 2 0 0 0-2 2v160a2 2 0 0 0 2 2h56a6 6 0 0 1 6 6Zm110.24-92.24l-40-40a6 6 0 0 0-8.48 8.48L201.51 122H104a6 6 0 0 0 0 12h97.51l-29.75 29.76a6 6 0 1 0 8.48 8.48l40-40a6 6 0 0 0 0-8.48Z"/>`),
		g.Group(children),
	)
}