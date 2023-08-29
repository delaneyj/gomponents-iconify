package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hospital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M16 20h4v3h-4v-3Zm4 5h-4v3h4v-3Zm-4 5h4v3h-4v-3Zm10-10h-4v3h4v-3Zm-4 5h4v3h-4v-3Zm4 5h-4v3h4v-3Zm2-10h4v3h-4v-3Zm4 5h-4v3h4v-3Zm-4 5h4v3h-4v-3Z"/><path fill-rule="evenodd" d="M31 6a2 2 0 0 0-2-2H19a2 2 0 0 0-2 2h-7v2h2v34H9a1 1 0 1 0 0 2h30a1 1 0 1 0 0-2h-3V8h2V6h-7Zm-2 12a2 2 0 0 0 2-2v-5h3v31h-3v-4h1v-2H16v2h1v4h-3V11h3v5a2 2 0 0 0 2 2h10Zm-6 24h2v-4h-2v4ZM34 8v1h-3V8h3ZM14 9h3V8h-3v1Zm6 1h3V7h2v3h3v2h-3v3h-2v-3h-3v-2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}