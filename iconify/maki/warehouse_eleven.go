package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarehouseEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M3 11H0V8h3v3zm4-3H4v3h3V8zm4 0H8v3h3V8zM5 4H2v3h3V4zm4 0H6v3h3V4zm1.44-.76a.5.5 0 0 0-.19-.68L5.75.06a.5.5 0 0 0-.49 0l-4.5 2.5a.5.5 0 0 0 .49.87L5.5 1.07l4.26 2.37a.5.5 0 0 0 .679-.198l.001-.002z" fill="currentColor"/>`),
		g.Group(children),
	)
}