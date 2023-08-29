package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpenPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M9.5 9.5H14a3.5 3.5 0 0 1 3.5 3.5v3a3.5 3.5 0 0 1-3.5 3.5H9A3.5 3.5 0 0 1 5.5 16v-3A3.502 3.502 0 0 1 8 9.645V6.75C8 4.693 9.552 3 11.5 3S15 4.693 15 6.75V7a.75.75 0 0 1-1.5 0v-.25c0-1.257-.91-2.25-2-2.25c-1.09 0-2 .993-2 2.25V9.5Z" clip-rule="evenodd" opacity=".8"/><path fill-rule="evenodd" d="M7.5 12.5a2 2 0 1 0 4 0a2 2 0 0 0-4 0Zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12 7.5H7A3.5 3.5 0 0 0 3.5 11v3A3.5 3.5 0 0 0 7 17.5h5a3.5 3.5 0 0 0 3.5-3.5v-3A3.5 3.5 0 0 0 12 7.5ZM4.5 11A2.5 2.5 0 0 1 7 8.5h5a2.5 2.5 0 0 1 2.5 2.5v3a2.5 2.5 0 0 1-2.5 2.5H7A2.5 2.5 0 0 1 4.5 14v-3Z" clip-rule="evenodd"/><path d="M7 8a.5.5 0 0 1-1 0V4.5a3.5 3.5 0 1 1 7 0v1a.5.5 0 0 1-1 0v-1a2.5 2.5 0 0 0-5 0V8Z"/></g>`),
		g.Group(children),
	)
}