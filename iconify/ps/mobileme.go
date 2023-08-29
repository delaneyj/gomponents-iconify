package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobileme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M462 257q0 38-26.5 64.5T372 348H92q-37 0-63.5-26.5T2 257q0-34 22-59.5T78 167q-1-3-1-10q0-24 17-41t40-17q20 0 37 14q17-37 39.5-57T276 36q55 0 84.5 36t29.5 88v8q31 6 51.5 31t20.5 58z"/>`),
		g.Group(children),
	)
}