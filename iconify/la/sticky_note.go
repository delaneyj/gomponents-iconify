package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickyNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h15.406l.313-.281l6-6l.281-.313V5zm2 2h18v12h-6v6H7zm14 14h2.563L21 23.563z"/>`),
		g.Group(children),
	)
}