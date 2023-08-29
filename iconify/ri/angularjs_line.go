package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngularjsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.523 16.65l.49-.27l1.118-9.71l-7.13-2.546l-7.132 2.545l1.119 9.71l.474.263L12 4.21l5.523 12.44Zm-1.098.61h-.799l-1.168-2.92H9.523l-1.169 2.92h-.778L12 19.713l4.425-2.453ZM12 2l9.3 3.32l-1.418 12.31L12 22l-7.88-4.37L2.7 5.32L12 2Zm1.698 10.54L12 8.45l-1.698 4.09h3.396Z"/>`),
		g.Group(children),
	)
}