package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5v15.574l-7 4.084l-6.074-3.544L2 21.5V5.926l7-4.084ZM4 7.074V18.5l5.074-2.114L15 19.842l5-2.916V5.5l-5.074 2.114L9 4.158L4 7.074ZM7.5 9.5a.25.25 0 1 0 0-.5a.25.25 0 0 0 0 .5Zm-1.75-.25a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0ZM14 10.5a.25.25 0 1 0 0-.5a.25.25 0 0 0 0 .5Zm-1.75-.25a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0ZM10 14a.25.25 0 1 0 0-.5a.25.25 0 0 0 0 .5Zm-1.75-.25a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0ZM16.5 15a.25.25 0 1 0 0-.5a.25.25 0 0 0 0 .5Zm-1.75-.25a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0Z"/>`),
		g.Group(children),
	)
}