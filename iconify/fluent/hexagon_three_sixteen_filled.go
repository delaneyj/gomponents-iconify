package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonThreeSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.473 1.998a1 1 0 0 1 .865-.498h2.324a1 1 0 0 1 .865.498l1.16 2a1 1 0 0 1 0 1.004l-1.16 2a1 1 0 0 1-.865.498H4.338a1 1 0 0 1-.865-.498l-1.16-2a1 1 0 0 1 0-1.004l1.16-2ZM4.338 8.5a1 1 0 0 0-.865.498l-1.16 2a1 1 0 0 0 0 1.004l1.16 2a1 1 0 0 0 .865.498h2.324a1 1 0 0 0 .865-.498l1.16-2a1 1 0 0 0 0-1.004l-1.16-2a1 1 0 0 0-.865-.498H4.338Zm5.135-3.002A1 1 0 0 1 10.338 5h2.324a1 1 0 0 1 .865.498l1.16 2a1 1 0 0 1 0 1.004l-1.16 2a1 1 0 0 1-.865.498h-2.324a1 1 0 0 1-.865-.498l-1.16-2a1 1 0 0 1 0-1.004l1.16-2Z"/>`),
		g.Group(children),
	)
}