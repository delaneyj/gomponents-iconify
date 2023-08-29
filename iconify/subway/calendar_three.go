package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M405.3 0h-42.7v64h42.7V0zm-256 0h-42.7v64h42.7V0zm277.4 0v85.3h-85.3V0H170.7v85.3H85.3V0H0v512h512V0h-85.3zm42.6 469.3H42.7V128h426.7v341.3zm-64-245.3l-42.7-42.7l-149.3 149.4l-64-64l-42.7 42.7L213.3 416l192-192z"/>`),
		g.Group(children),
	)
}