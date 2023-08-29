package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Awards(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19 9a6.992 6.992 0 0 1-3 5.745V22h-2.586L12 20.586L10.586 22H8v-7.255A7 7 0 1 1 19 9Zm-2 0A5 5 0 1 1 7 9a5 5 0 0 1 10 0Zm-7 10.757l2-2l2 2V15.71a7.001 7.001 0 0 1-2 .29a7.001 7.001 0 0 1-2-.29v4.047Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}