package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldDismissSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.647 2.146a.5.5 0 0 1 .708 0C9.595 3.39 10.969 4 12.5 4a.5.5 0 0 1 .5.5v3.001c0 3.219-1.641 5.407-4.842 6.473a.499.499 0 0 1-.316 0C4.642 12.908 3 10.72 3 7.501V4.5a.5.5 0 0 1 .5-.5c1.53 0 2.904-.611 4.147-1.854ZM6.534 5.84a.5.5 0 0 0-.638.057l-.057.07a.5.5 0 0 0 .057.638L7.293 8L5.896 9.396l-.057.07a.5.5 0 0 0 .057.638l.07.057a.5.5 0 0 0 .638-.057L8 8.707l1.396 1.397l.07.057a.5.5 0 0 0 .638-.057l.057-.07a.5.5 0 0 0-.057-.638L8.707 8l1.397-1.396l.057-.07a.5.5 0 0 0-.057-.638l-.07-.057a.5.5 0 0 0-.638.057L8 7.293L6.604 5.896l-.07-.057Z"/>`),
		g.Group(children),
	)
}