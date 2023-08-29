package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m273.637 128l3.809 32h96.43l-8 64H252.573l-7.442-64h.089l-3.809-32l-.045-.389l-9.041-77.745L230.247 32H104v32h97.753l7.442 64h-91.319l40.5 323.969A32.051 32.051 0 0 0 190.125 480h147.75a32.051 32.051 0 0 0 31.753-28.031L410.124 128Zm-119.513 32h58.792l7.442 64h-58.234Zm183.765 288H190.124l-24-192h57.955l13.953 120h32.215l-13.954-120H361.88Z"/>`),
		g.Group(children),
	)
}