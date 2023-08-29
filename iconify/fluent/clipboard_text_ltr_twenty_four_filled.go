package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardTextLtrTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.75 3.5h-3.5a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5Zm0-1.5c1.14 0 2.08.846 2.23 1.945l.013.135l-.007-.08h1.764A2.25 2.25 0 0 1 20 6.25v13.5A2.25 2.25 0 0 1 17.75 22H6.25A2.25 2.25 0 0 1 4 19.75V6.25A2.25 2.25 0 0 1 6.25 4h1.764l-.008.08l.015-.135A2.25 2.25 0 0 1 10.25 2h3.5ZM14 17H8a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 0-1.5Zm-2-4H8a.75.75 0 0 0 0 1.5h4a.75.75 0 0 0 0-1.5Zm4-4H8a.75.75 0 0 0 0 1.5h8A.75.75 0 0 0 16 9Z"/>`),
		g.Group(children),
	)
}