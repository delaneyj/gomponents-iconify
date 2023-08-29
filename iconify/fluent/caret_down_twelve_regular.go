package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDownTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.076 4.617A1 1 0 0 1 4 4h4a1 1 0 0 1 .707 1.707l-2 2a1 1 0 0 1-1.414 0l-2-2a1 1 0 0 1-.217-1.09ZM4 5l1.823 1.823a.25.25 0 0 0 .354 0L8 5H4Z"/>`),
		g.Group(children),
	)
}