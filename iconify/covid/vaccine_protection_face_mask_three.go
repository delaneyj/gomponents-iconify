package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VaccineProtectionFaceMaskThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 11.127a38.599 38.599 0 0 1 5-.5m-5 4.125c1.65.31 3.322.477 5 .5"/><path d="M12 19.5c2.743.074 5.47-.437 8-1.5V9s-4.9-4.5-8-4.5C8.9 4.5 4 9 4 9v9a19.313 19.313 0 0 0 8 1.5Z"/><path d="M17 11.127a38.598 38.598 0 0 0-5-.5m5 4.125c-1.65.31-3.322.477-5 .5M4 18a6.828 6.828 0 0 1-3.25-5.814V8.054a1.082 1.082 0 0 1 1.166-1.08A1.467 1.467 0 0 1 3 7.59L4 9m16 9a6.828 6.828 0 0 0 3.25-5.816v-4.13a1.083 1.083 0 0 0-1.166-1.08A1.467 1.467 0 0 0 21 7.59L20 9"/></g>`),
		g.Group(children),
	)
}