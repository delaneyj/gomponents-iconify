package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M21.264 2.293a1 1 0 0 0-1.415 0l-.872.872a3.001 3.001 0 0 0-3.415.587L4.955 14.358l5.657 5.657L21.22 9.408a2.999 2.999 0 0 0 .586-3.414l.873-.873a1 1 0 0 0 0-1.414l-1.415-1.414Zm-4.268 8.51l-6.384 6.384l-2.828-2.829l6.383-6.383l2.829 2.829Zm1.818-1.818l.99-.99a1 1 0 0 0 0-1.415l-1.413-1.414a1 1 0 0 0-1.415 0l-.99.99l2.828 2.83Z" clip-rule="evenodd"/><path d="m2 22.95l2.122-7.778l5.656 5.657L2 22.95Z"/></g>`),
		g.Group(children),
	)
}