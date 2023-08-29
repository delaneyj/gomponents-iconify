package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5v15.574l-7 4.084l-6.074-3.544L2 21.5V5.926l7-4.084ZM4 7.074V18.5l5.074-2.114L15 19.842l5-2.916V5.5l-5.074 2.114L9 4.158L4 7.074Z"/>`),
		g.Group(children),
	)
}