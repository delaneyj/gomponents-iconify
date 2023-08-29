package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M36.08 43.37A6 6 0 0 0 34 47.9V216a6 6 0 0 0 12 0v-45.22c28.08-22.79 51.88-11 79.34 2.57c16.12 8 33.49 16.58 52 16.58c13.57 0 27.76-4.6 42.56-17.42A6 6 0 0 0 222 168V47.9a6 6 0 0 0-9.93-4.54c-29 25.12-53.28 13.09-81.41-.84c-27.89-13.81-59.66-29.36-94.58.85ZM210 165.17c-28.08 22.8-51.88 11-79.34-2.58C105.4 150.08 77.09 136.07 46 156V50.72c28.08-22.8 51.88-11 79.34 2.56C150.6 65.79 178.91 79.8 210 59.91Z"/>`),
		g.Group(children),
	)
}