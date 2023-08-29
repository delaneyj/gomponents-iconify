package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GpsBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M2 12c0 .385.312.698.698.698H4.59a7.444 7.444 0 0 0 6.712 6.712v1.892a.698.698 0 0 0 1.396 0V19.41a7.444 7.444 0 0 0 6.712-6.712h1.892a.698.698 0 0 0 0-1.396H19.41a7.444 7.444 0 0 0-6.712-6.712V2.698a.698.698 0 0 0-1.396 0V4.59a7.444 7.444 0 0 0-6.712 6.712H2.698A.698.698 0 0 0 2 12Zm6.512 0a3.488 3.488 0 1 1 6.976 0a3.488 3.488 0 0 1-6.976 0Z" clip-rule="evenodd"/><path d="M9.907 12a2.093 2.093 0 1 1 4.186 0a2.093 2.093 0 0 1-4.186 0Z"/></g>`),
		g.Group(children),
	)
}