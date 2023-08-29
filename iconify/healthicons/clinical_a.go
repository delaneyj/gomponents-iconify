package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClinicalA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19 4a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h1v2.273a1.5 1.5 0 0 0-1 1.415V39a5 5 0 0 0 10 0V12.687a1.5 1.5 0 0 0-1-1.414V9h1a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H19Zm3 7.688V9h4v2.688a1.5 1.5 0 0 0 1 1.414V18h-6v-4.898a1.5 1.5 0 0 0 1-1.415Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}