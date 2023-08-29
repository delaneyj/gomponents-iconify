package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Discount(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.496 1.19l10.506.808l.808 10.505l-10.932 10.932L.564 12.121L11.496 1.19Zm.764 2.064l-8.867 8.867l8.485 8.486l8.867-8.868l-.606-7.879l-7.879-.606Zm3.86 4.624a1 1 0 1 0-1.414 1.414a1 1 0 0 0 1.414-1.414Zm-2.828-1.414a3 3 0 1 1 4.243 4.242a3 3 0 0 1-4.243-4.242Z"/>`),
		g.Group(children),
	)
}