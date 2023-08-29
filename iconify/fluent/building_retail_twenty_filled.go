package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingRetailTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 10V8h8v2H6Zm6 4v-1h2v1h-2Zm6-9.5a1.5 1.5 0 0 1-1 1.415V16.5a.5.5 0 0 1-.5.5H10v-4.5a.5.5 0 0 0-.5-.5h-4a.5.5 0 0 0-.5.5V17H3.5a.5.5 0 0 1-.5-.5V5.915A1.5 1.5 0 0 1 3.5 3h13A1.5 1.5 0 0 1 18 4.5Zm-15 0a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 0-1h-13a.5.5 0 0 0-.5.5ZM5.5 7a.5.5 0 0 0-.5.5v3a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-.5-.5h-9Zm6 5a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5h-3ZM6 17h3v-4H6v4Z"/>`),
		g.Group(children),
	)
}