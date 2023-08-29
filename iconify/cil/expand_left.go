package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M464 16h32v480h-32zm-320 0H16v480h128Zm-32 448H48V48h64ZM416 96h-37.86L208.776 255.923L376.079 416H416Zm-32 283.291L255.224 256.077L384 134.478Z"/>`),
		g.Group(children),
	)
}