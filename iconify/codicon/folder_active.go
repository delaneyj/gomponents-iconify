package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderActive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M7.71 3h6.79l.51.5v4.507A4.997 4.997 0 0 0 14 7.416V5.99H7.69l-.86.86l-.35.15H1.99v6H7.1c.07.348.177.682.316 1H1.51l-.5-.5v-11l.5-.5h5l.35.15l.85.85zm-.22 2h6.5l.01-.99H7.5l-.36-.15l-.85-.85H2v3h4.28l.86-.86l.35-.15z"/><path d="M9.778 8.674a4 4 0 1 1 4.444 6.652a4 4 0 0 1-4.444-6.652zm2.13 4.99l2.387-3.182l-.8-.6l-2.077 2.769l-1.301-1.041l-.625.78l1.704 1.364l.713-.09z"/></g>`),
		g.Group(children),
	)
}