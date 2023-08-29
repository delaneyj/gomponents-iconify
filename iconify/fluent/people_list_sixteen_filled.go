package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeopleListSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 5.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0ZM11.5 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM8 9.5c0-.174.03-.342.085-.498A1.528 1.528 0 0 0 8 9H3a1.5 1.5 0 0 0-1.5 1.5v.075s0 2.925 4 2.925c1.21 0 2.055-.268 2.644-.642c.062-.13.142-.251.238-.358a1.494 1.494 0 0 1-.382-1c0-.384.144-.735.382-1A1.494 1.494 0 0 1 8 9.5ZM9.5 9a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1h-5Zm0 2a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1h-5Zm0 2a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1h-5Z"/>`),
		g.Group(children),
	)
}