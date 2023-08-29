package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20Zm83.13 96h-31.57a155.5 155.5 0 0 0-19.51-65.63A84.23 84.23 0 0 1 211.13 116Zm-110.62 24h55c-2.25 26.69-12 51.46-27.49 69.85c-15.56-18.39-25.26-43.16-27.51-69.85Zm0-24c2.25-26.69 11.95-51.46 27.49-69.85c15.54 18.39 25.24 43.16 27.49 69.85ZM96 50.37A155.5 155.5 0 0 0 76.44 116H44.87A84.23 84.23 0 0 1 96 50.37ZM44.87 140h31.57A155.5 155.5 0 0 0 96 205.63A84.23 84.23 0 0 1 44.87 140Zm115.18 65.63A155.5 155.5 0 0 0 179.56 140h31.57a84.23 84.23 0 0 1-51.08 65.63Z"/>`),
		g.Group(children),
	)
}