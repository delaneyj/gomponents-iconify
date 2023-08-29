package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func INoteActionNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsINoteActionNegative0)"><path d="M14 19a1 1 0 0 1 1-1h13a1 1 0 1 1 0 2H15a1 1 0 0 1-1-1Zm1 4a1 1 0 1 0 0 2h7a1 1 0 1 0 0-2h-7Zm-1 6a1 1 0 0 1 1-1h11a1 1 0 1 1 0 2H15a1 1 0 0 1-1-1Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM15 8h-2a3 3 0 0 0-3 3v24a3 3 0 0 0 3 3h17a3 3 0 0 0 3-3V11a3 3 0 0 0-3-3h-1v2a1 1 0 1 1-2 0V7a1 1 0 1 0-2 0v1h-6v2a1 1 0 1 1-2 0V7a1 1 0 1 0-2 0v1Zm-7 2v25a5 5 0 0 0 5 5h17v2H13a7 7 0 0 1-7-7V10h2Zm28 3a3 3 0 1 1 6 0v3h-6v-3Zm0 20V18h6v15l-3 4l-3-4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsINoteActionNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}