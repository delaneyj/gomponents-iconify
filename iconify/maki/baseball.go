package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Baseball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M9 3.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0zM6 .28A.28.28 0 0 0 5.72 0a.49.49 0 0 0-.25.16L3 4.59a.48.48 0 0 0 0 .13c0 .155.125.28.28.28a.49.49 0 0 0 .26-.16L6 .41a.472.472 0 0 0 0-.13zm5.9 13.92L9 6.39A.49.49 0 0 0 8.52 6h-5a.5.5 0 0 0 0 1H6l1.45 2.51l-4.27 4.61a.49.49 0 0 0-.18.38a.5.5 0 0 0 .5.5a.49.49 0 0 0 .33-.13l4.45-4.15l2.76 4a.51.51 0 0 0 .44.26c.28 0 .51-.22.52-.5a.5.5 0 0 0-.1-.28z"/>`),
		g.Group(children),
	)
}