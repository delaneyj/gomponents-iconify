package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M35 9.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Z"/><path fill-rule="evenodd" d="M31.5 11a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 2a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7ZM19 16a2 2 0 1 0 0 4h7v20a2 2 0 1 0 4 0v-9h3v9a2 2 0 1 0 4 0V27.917A6.002 6.002 0 0 0 36 16H19Zm20 6a2 2 0 0 0-2-2v4a2 2 0 0 0 2-2Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6 32a2 2 0 0 1 2-2h13a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2v-8Zm4 7v-6h2v6h-2Zm7-6v6h2v-6h-2Zm-6-5a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v2h-2v-2h-3v2h-2v-2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}