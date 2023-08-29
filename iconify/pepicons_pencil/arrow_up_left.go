package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5.948 12.145a.5.5 0 0 1-.453-.543l.471-5.186a.5.5 0 0 1 .996.09l-.471 5.186a.5.5 0 0 1-.543.453Z"/><path d="M12.148 5.945a.5.5 0 0 1-.453.543L6.51 6.96a.5.5 0 0 1-.09-.996l5.185-.472a.5.5 0 0 1 .543.453Z"/><path d="M6.647 6.643a.5.5 0 0 1 .707 0l6.535 6.536a.5.5 0 1 1-.707.707L6.646 7.351a.5.5 0 0 1 0-.708Z"/></g>`),
		g.Group(children),
	)
}