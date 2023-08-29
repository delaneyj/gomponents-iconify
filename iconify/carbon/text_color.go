package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 21h2L17 4h-2L8 21h2l1.61-4h8.74zm-9.57-6l3.44-8.37h.26L19.54 15zM6 24h20v4H6z"/>`),
		g.Group(children),
	)
}