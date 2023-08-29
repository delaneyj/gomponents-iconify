package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresentationLineLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3h-3.586l2.293 2.293a1 1 0 0 1-1.414 1.414L12 17.414l-3.293 3.293a1 1 0 0 1-1.414-1.414L9.586 17H6a3 3 0 0 1-3-3V6zm3-1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H6zm10.6 2.2a1 1 0 0 1 .2 1.4l-3 4a1 1 0 0 1-1.507.107l-2.138-2.137l-1.323 1.985a1 1 0 0 1-1.664-1.11l2-3a1 1 0 0 1 1.54-.152l2.184 2.185L15.2 7.4a1 1 0 0 1 1.4-.2z"/></g>`),
		g.Group(children),
	)
}