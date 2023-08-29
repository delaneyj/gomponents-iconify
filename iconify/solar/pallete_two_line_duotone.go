package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PalleteTwoLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 12.026c0 5.146 3.867 9.387 8.847 9.96c.735.085 1.447-.228 1.97-.753a1.68 1.68 0 0 0 0-2.372c-.523-.525-.95-1.307-.555-1.934c1.576-2.508 9.738 3.251 9.738-4.9C22 6.488 17.523 2 12 2S2 6.489 2 12.026Z"/><circle cx="17.5" cy="11.5" r="1.5" opacity=".5"/><circle cx="6.5" cy="11.5" r="1.5" opacity=".5"/><path d="M11 7a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm5 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}