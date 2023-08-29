package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M149.3 0h-42.7v64h42.7V0zm256 0h-42.7v64h42.7V0zM148.9 426.7h214.3c33 0 52.3-22.2 35.9-49.4L286 190.9c-16.5-27.2-43.4-27.2-59.9 0L113 377.3c-16.5 27.1 2.9 49.4 35.9 49.4zm85.8-192h42.7v106.7h-42.7V234.7zm0 128h42.7v42.7h-42.7v-42.7zM426.7 0v85.3h-85.3V0H170.7v85.3H85.3V0H0v512h512V0h-85.3zm42.6 469.3H42.7V128h426.7v341.3z"/>`),
		g.Group(children),
	)
}