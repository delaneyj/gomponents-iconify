package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryChargeMinimalisticLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 12c0-3.771 0-5.657 1.172-6.828C4.343 4 6.229 4 10 4h1.5c3.771 0 5.657 0 6.828 1.172C19.5 6.343 19.5 8.229 19.5 12c0 3.771 0 5.657-1.172 6.828C17.157 20 15.271 20 11.5 20H10c-3.771 0-5.657 0-6.828-1.172C2 17.657 2 15.771 2 12Z"/><path stroke-linecap="round" d="M22 14v-4"/><path stroke-linecap="round" stroke-linejoin="round" d="M11.5 9L9 12h3.5L10 15"/></g>`),
		g.Group(children),
	)
}