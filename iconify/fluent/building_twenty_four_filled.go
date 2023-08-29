package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 4.25A2.25 2.25 0 0 1 6.25 2h8a2.25 2.25 0 0 1 2.25 2.25V9.5h1.25A2.25 2.25 0 0 1 20 11.75v9.5a.75.75 0 0 1-.75.75H16.5v-3.75c0-.69-.56-1.25-1.25-1.25h-6.5c-.69 0-1.25.56-1.25 1.25V22H4.75a.75.75 0 0 1-.75-.75v-17ZM15 18.5V22h-2.25v-3.5H15Zm-3.75 0V22H9v-3.5h2.25ZM7.5 6.5a1 1 0 1 0 2 0a1 1 0 0 0-2 0Zm1 6a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm0-3.5a1 1 0 1 0 0 2a1 1 0 0 0 0-2ZM12 5.5a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm0 7a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm3.5 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2ZM12 9a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}