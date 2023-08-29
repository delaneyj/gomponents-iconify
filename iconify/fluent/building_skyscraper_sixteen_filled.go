package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingSkyscraperSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.75 1h-.5a.25.25 0 0 0-.25.25V2h1v-.75A.25.25 0 0 0 6.75 1ZM8.5 3h-4a.5.5 0 0 0-.5.5V4h5v-.5a.5.5 0 0 0-.5-.5Zm3 9a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm.5-2.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM3 6.45A1.7 1.7 0 0 1 3.812 5h5.376A1.7 1.7 0 0 1 10 6.45v.05h.3a3.7 3.7 0 0 1 3.7 3.7v3.3a.5.5 0 0 1-.5.5H7v-2.1a.4.4 0 0 0-.4-.4h-.2a.4.4 0 0 0-.4.4V14H3.5a.5.5 0 0 1-.5-.5V6.45Zm10 3.75a2.7 2.7 0 0 0-2.7-2.7H10V13h3v-2.8ZM6 7.5a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0Zm1.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM6 9.5a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0Zm1.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/>`),
		g.Group(children),
	)
}