package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleMusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.86a18.79 18.79 0 0 0-12.17 4.44c-5.72 4.79-6.95 12.77-6.12 19.51L7 40.34a1.73 1.73 0 0 0 1.93 1.5l5.07-.63a1.72 1.72 0 0 0 1.5-1.92l-1.3-10.88A1.67 1.67 0 0 0 12.29 27l-1.85.26a15 15 0 0 1-.22-2.42a13.81 13.81 0 0 1 27.62 0a13.4 13.4 0 0 1-.22 2.37l-1.82-.22a1.72 1.72 0 0 0-1.92 1.5l-1.34 10.83A1.72 1.72 0 0 0 34 41.24l5.06.62a1.72 1.72 0 0 0 1.94-1.5l1.21-10.53c.87-6.73-.43-14.75-6.12-19.5a18.65 18.65 0 0 0-12.15-4.49Z"/>`),
		g.Group(children),
	)
}