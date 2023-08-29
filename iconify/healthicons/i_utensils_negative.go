package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IUtensilsNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsIUtensilsNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM19 4h9v4h-4v2h4v3h-2v2h2v3h-2v2h2v3h-4v2h4v3h-2v2h2v3h-2v2h2v3h-4v2h4v4h-9V4Zm14 24a2 2 0 0 0-2 2v5h7v-5a2 2 0 0 0-2-2h-3Zm5 9h-7v5a2 2 0 0 0 2 2h3a2 2 0 0 0 2-2v-5ZM10 14a3 3 0 1 1 6 0v3h-6v-3Zm0 26V19h6v21l-3 4l-3-4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsIUtensilsNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}