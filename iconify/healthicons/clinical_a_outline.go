package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClinicalAOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19 4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h1v1.273a1.5 1.5 0 0 0-1 1.415V39a5 5 0 0 0 10 0V12.687a1.5 1.5 0 0 0-1-1.414V10h1a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H19Zm2 9.102a1.5 1.5 0 0 0 1-1.415V10h4v1.688a1.5 1.5 0 0 0 1 1.414V18h-6v-4.898Zm6 11.883L22.015 20H21v1.157l6 6v-2.172ZM24.844 20L27 22.157v-1.172L26.015 20h-1.171ZM27 29.985l-6-6v2.172l6 6v-2.172Zm0 7.172l-6-6v-2.172l6 6v2.172Zm-.125 2.703L21 33.985v2.172l5.041 5.041c.387-.359.679-.819.834-1.338ZM24 42a3 3 0 0 1-3-3v-.015L24.015 42H24ZM20 6v2h8V6h-8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}