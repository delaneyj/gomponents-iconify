package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backspace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m17.743 8.464l1.414 1.415L17.036 12l2.121 2.121l-1.414 1.415l-2.122-2.122l-2.121 2.122l-1.414-1.415L14.207 12l-2.121-2.121L13.5 8.465l2.121 2.12l2.122-2.12Z"/><path fill-rule="evenodd" d="m8.586 19l-6.293-6.293a1 1 0 0 1 0-1.414L8.586 5h14v14h-14Zm.828-12l-5 5l5 5h11.172V7H9.414Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}