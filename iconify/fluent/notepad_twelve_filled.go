package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotepadTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 1.5a.5.5 0 0 1 1 0V2h2v-.5a.5.5 0 0 1 1 0V2a2 2 0 0 1 2 2v3H7.5A1.5 1.5 0 0 0 6 8.5V11H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2v-.5Zm0 3a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 0-1h-3a.5.5 0 0 0-.5.5Zm0 2a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 0-1h-1a.5.5 0 0 0-.5.5ZM4.5 9a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM7 8.5a.5.5 0 0 1 .5-.5H10l-3 3V8.5Z"/>`),
		g.Group(children),
	)
}