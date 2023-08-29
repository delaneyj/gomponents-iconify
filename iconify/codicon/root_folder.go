package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RootFolder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M7.71 3h6.79l.51.5v10l-.5.5H8.743a5.48 5.48 0 0 0 .657-1h4.59v-1.51l.01-4v-1.5H7.69l-.017.017a5.494 5.494 0 0 0-.881-.508l.348-.349l.35-.15h6.5l.01-.99H7.5l-.36-.15l-.85-.85H2V5.6a5.45 5.45 0 0 0-.99.649V2.5l.5-.5h5l.35.15l.85.85z" clip-rule="evenodd"/><path d="M6 10.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0z"/><path fill-rule="evenodd" d="M8 10.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0zM4.5 13a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}