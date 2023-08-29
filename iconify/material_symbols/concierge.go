package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Concierge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 22v-2h13v2H10Zm1-3q0-2.025 1.275-3.538T15.5 13.6v-.625q0-.425.288-.712t.712-.288q.425 0 .713.288t.287.712v.625q1.925.35 3.213 1.863T22 19H11ZM1 13h4V2H1v11Zm6 0h1.975l7.05-2.625q.425-.175.7-.563T17 8.95q0-.4-.275-.675T16.05 8H13l-1.75.675l-.35-.925L13 7h7q.8 0 1.4-.575T22 5l-8-3l-7 1.95V13Z"/>`),
		g.Group(children),
	)
}