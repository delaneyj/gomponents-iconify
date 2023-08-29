package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BowlingOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.75a9.25 9.25 0 1 0 0 18.5a9.25 9.25 0 0 0 0-18.5ZM1.25 12C1.25 6.063 6.063 1.25 12 1.25S22.75 6.063 22.75 12S17.937 22.75 12 22.75S1.25 17.937 1.25 12ZM12 6.25a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5ZM9.75 7a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0ZM8 8.75a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5Zm-2.25.75a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0ZM12 11.25a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5ZM9.75 12a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}