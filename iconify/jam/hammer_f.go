package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HammerF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.288 5.703a1.5 1.5 0 0 1 .293-1.707L3.996 2.58a1.5 1.5 0 0 1 1.707-.293L7.178.814a2 2 0 0 1 2.828 0l2.829 2.828a2 2 0 0 1 0 2.829L6.47 12.835a2 2 0 0 1-2.829 0L.814 10.006a2 2 0 0 1 0-2.828l1.474-1.475zM9.3 12.835l3.536-3.536l3.535 3.536l-3.535 3.535L9.3 12.835zm4.95 4.95l3.535-3.536l1.768 1.768a2.5 2.5 0 0 1-3.535 3.535l-1.768-1.768z"/>`),
		g.Group(children),
	)
}