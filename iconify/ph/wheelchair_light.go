package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WheelchairLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M253.69 190.1a6 6 0 0 0-7.59-3.79l-19.1 6.35l-29.68-59.34A6 6 0 0 0 192 130h-82v-27.95h58a6 6 0 0 0 0-12h-58V77.4a30 30 0 1 0-12 0v14A70 70 0 0 0 112 230c32.62 0 62-22.08 69.81-52.5a6 6 0 0 0-11.62-3c-6.49 25.21-31 43.5-58.19 43.5a58 58 0 0 1-14-114.28V136a6 6 0 0 0 6 6h84.29l30.34 60.68a6 6 0 0 0 7.27 3l24-8a6 6 0 0 0 3.79-7.58ZM86 48a18 18 0 1 1 18 18a18 18 0 0 1-18-18Z"/>`),
		g.Group(children),
	)
}