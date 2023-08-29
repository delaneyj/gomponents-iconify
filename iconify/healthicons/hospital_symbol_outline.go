package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalSymbolOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24 42c-9.941 0-18-8.059-18-18S14.059 6 24 6s18 8.059 18 18s-8.059 18-18 18ZM4 24C4 12.954 12.954 4 24 4s20 8.954 20 20s-8.954 20-20 20S4 35.046 4 24Zm15.991 9.048v-7h8v7h4v-18h-4v7h-8v-7h-4v18h4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}