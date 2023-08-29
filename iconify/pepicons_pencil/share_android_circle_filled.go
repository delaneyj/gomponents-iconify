package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareAndroidCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilShareAndroidCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M8 15.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0-4a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm9-1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0-4a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm0 14a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0-4a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Z" clip-rule="evenodd"/><path d="m9.754 12.18l-.508-.86l5.5-3.25l.508.86l-5.5 3.25ZM15 17.878l.479-.878l-5.5-3l-.479.878l5.5 3Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilShareAndroidCircleFilled0)"/></g>`),
		g.Group(children),
	)
}