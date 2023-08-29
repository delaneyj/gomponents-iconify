package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hamburger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.706 4.073C7.366 2.7 9.605 2 12 2c2.396 0 4.635.7 6.294 2.073C19.974 5.462 21 7.495 21 10v1H3v-1c0-2.505 1.027-4.538 2.706-5.927ZM5.074 9h13.852c-.219-1.432-.911-2.563-1.907-3.386C15.784 4.59 14.023 4 12 4s-3.784.591-5.02 1.614C5.986 6.437 5.294 7.568 5.075 9ZM6 11.798l3 2l3-2l3 2l3-2l4.387 2.925l-1.11 1.664L18 14.202l-3 2l-3-2l-3 2l-3-2l-3.277 2.185l-1.11-1.664L6 11.798ZM3 17h18v1a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5v-1Zm2.17 2A3.001 3.001 0 0 0 8 21h8a3.001 3.001 0 0 0 2.83-2H5.17Z"/>`),
		g.Group(children),
	)
}