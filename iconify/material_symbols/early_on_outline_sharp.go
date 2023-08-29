package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EarlyOnOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 21q-1.25 0-2.125-.875T14 18q0-1.25.875-2.125T17 15q1.25 0 2.125.875T20 18q0 1.25-.875 2.125T17 21Zm-.75-7v-2h1.5v2h-1.5Zm0 10v-2h1.5v2h-1.5Zm4.125-8.325l-1.075-1.05l1.425-1.425l1.05 1.075l-1.4 1.4Zm-7.1 7.1l-1.05-1.05L13.65 20.3l1.05 1.05l-1.425 1.425ZM21 18.75v-1.5h2v1.5h-2Zm-10 0v-1.5h2v1.5h-2Zm9.725 4.025l-1.4-1.425l1.05-1.05l1.425 1.4l-1.075 1.075Zm-7.1-7.075l-1.4-1.425l1.05-1.05l1.425 1.4l-1.075 1.075ZM3 22V4h3V2h2v2h8V2h2v2h3v6H5v10h4v2H3ZM5 8h14V6H5v2Zm0 0V6v2Z"/>`),
		g.Group(children),
	)
}