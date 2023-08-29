package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M149.3 0h-42.7v64h42.7V0zm256 0h-42.7v64h42.7V0zm21.4 0v85.3h-85.3V0H170.7v85.3H85.3V0H0v512h512V0h-85.3zm42.6 469.3H42.7V128h426.7v341.3zm-234.6-64h42.7V320h85.3v-42.7h-85.3V192h-42.7v85.3h-85.3V320h85.3v85.3z"/>`),
		g.Group(children),
	)
}