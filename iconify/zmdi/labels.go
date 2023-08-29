package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Labels(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M26 363q-16-7-22.5-23T3 308l52-125v192zm416-79q7 16 0 32.5T419 340l-157 65q-8 3-16 3q-29 0-39-26L101 126q-4-8-3-17q0-28 26-38L281 6q8-3 17-3q28 0 39 26zM140.5 131q8.5 0 15-6.5t6.5-15t-6.5-15t-15-6.5t-15 6.5t-6.5 15t6.5 15t15 6.5zM98 365V230l73 178h-31q-17 0-29.5-12.5T98 365z"/>`),
		g.Group(children),
	)
}