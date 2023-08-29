package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func INoteAction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M26 6a1 1 0 0 1 1 1v3a1 1 0 1 0 2 0V8h1a3 3 0 0 1 3 3v24a3 3 0 0 1-3 3H13a3 3 0 0 1-3-3V11a3 3 0 0 1 3-3h2V7a1 1 0 1 1 2 0v3a1 1 0 1 0 2 0V8h6V7a1 1 0 0 1 1-1ZM15 18a1 1 0 1 0 0 2h13a1 1 0 1 0 0-2H15Zm-1 6a1 1 0 0 1 1-1h7a1 1 0 1 1 0 2h-7a1 1 0 0 1-1-1Zm1 4a1 1 0 1 0 0 2h11a1 1 0 1 0 0-2H15Z" clip-rule="evenodd"/><path d="M13 42a7 7 0 0 1-7-7V10h2v25a5 5 0 0 0 5 5h17v2H13Zm23-29a3 3 0 1 1 6 0v3h-6v-3Zm0 20V18h6v15l-3 4l-3-4Z"/></g>`),
		g.Group(children),
	)
}