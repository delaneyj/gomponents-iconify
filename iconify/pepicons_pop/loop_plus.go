package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoopPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M4.828 4.828a6 6 0 1 0 8.486 8.486a6 6 0 0 0-8.486-8.486ZM11.9 11.9a4 4 0 1 1-5.656-5.656A4 4 0 0 1 11.9 11.9Z" clip-rule="evenodd"/><path d="M11.9 14.728a1.5 1.5 0 1 1 2.12-2.121l3.536 3.535a1.5 1.5 0 1 1-2.121 2.121l-3.536-3.535Zm-4.693-4.564a1 1 0 0 1 0-2h4a1 1 0 1 1 0 2h-4Z"/><path d="M8.207 7.164a1 1 0 0 1 2 0v4a1 1 0 1 1-2 0v-4Z"/></g>`),
		g.Group(children),
	)
}