package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2.5 11.25h2v1.25h-2z"/><path fill="currentColor" d="M14.75 4H13.3a2.1 2.1 0 0 0-1.87-3a4.24 4.24 0 0 0-2.84 1a3.42 3.42 0 0 0-.59.67A3.42 3.42 0 0 0 7.41 2a4.24 4.24 0 0 0-2.84-1A2.1 2.1 0 0 0 2.7 4H1.25A1.25 1.25 0 0 0 0 5.25v8.5A1.25 1.25 0 0 0 1.25 15h13.5A1.25 1.25 0 0 0 16 13.75v-8.5A1.25 1.25 0 0 0 14.75 4zM9.42 2.94a3 3 0 0 1 2-.69a.85.85 0 0 1 0 1.7H8.74a2.32 2.32 0 0 1 .68-1.01zm-4.85-.69a3 3 0 0 1 2 .69a2.32 2.32 0 0 1 .68 1H4.57a.85.85 0 0 1 0-1.7zm-3.32 3h13.5V6.5H1.25V5.25zm0 8.5V9h13.5v4.75z"/>`),
		g.Group(children),
	)
}