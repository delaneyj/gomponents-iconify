package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KissingFaceWithSmilingEyesFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm6.707-1.293a1 1 0 0 1-1.414-1.414l1-1a1 1 0 0 1 1.414 0l1 1a1 1 0 0 1-1.414 1.414L9 10.414l-.293.293Zm6.586 0a1 1 0 0 0 1.414-1.414l-1-1a1 1 0 0 0-1.414 0l-1 1a1 1 0 0 0 1.414 1.414l.293-.293l.293.293Zm-1.399 2.846a1 1 0 0 1-.447 1.341l-.21.106l.21.106a1 1 0 1 1-.894 1.788l-2-1a1 1 0 0 1 0-1.788l2-1a1 1 0 0 1 1.341.447Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}