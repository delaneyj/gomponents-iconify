package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InkingToolAccentThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 4h26v4a1 1 0 0 1-1 1h-5c0 .15-.034.303-.106.447l-5 10a1 1 0 1 1-1.788-.894L20.882 9H4a1 1 0 0 1-1-1V4Zm14.857 21.875c0 1.864-.863 3.375-1.928 3.375c-1.066 0-1.929-1.511-1.929-3.375s.864-3.375 1.929-3.375s1.928 1.511 1.928 3.375Z"/>`),
		g.Group(children),
	)
}