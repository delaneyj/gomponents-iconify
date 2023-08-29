package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlaneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236.2 218.3a15.87 15.87 0 0 1-17.54 4.76l-79.95-27.36a4 4 0 0 1-2.71-3.79V120a8 8 0 0 0-8.53-8a8.19 8.19 0 0 0-7.47 8.26v71.57a4 4 0 0 1-2.7 3.79l-80 27.44a16 16 0 0 1-19.25-22.92L114 32.13a16 16 0 0 1 27.89 0l96.01 167.96a15.89 15.89 0 0 1-1.7 18.21Z"/>`),
		g.Group(children),
	)
}