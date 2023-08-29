package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoopMinusCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopLoopMinusCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M7.828 7.828a6 6 0 1 0 8.486 8.486a6 6 0 0 0-8.486-8.486ZM14.9 14.9a4 4 0 1 1-5.656-5.656A4 4 0 0 1 14.9 14.9Z" clip-rule="evenodd"/><path d="M14.9 17.728a1.5 1.5 0 1 1 2.12-2.121l3.536 3.535a1.5 1.5 0 1 1-2.121 2.121L14.9 17.728Zm-5.193-4.814a1 1 0 1 1 0-2h4.485a1 1 0 1 1 0 2H9.707Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopLoopMinusCircleFilled0)"/></g>`),
		g.Group(children),
	)
}