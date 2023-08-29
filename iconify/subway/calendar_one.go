package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M362.7 320H320v42.7h42.7V320zm-85.4-170.7h-42.7V192h42.7v-42.7zm85.4 256H320V448h42.7v-42.7zm0-256H320V192h42.7v-42.7zM149.3 0h-42.7v64h42.7V0zM448 149.3h-42.7V192H448v-42.7zm-170.7 85.4h-42.7v42.7h42.7v-42.7zm170.7 0h-42.7v42.7H448v-42.7zm0 85.3h-42.7v42.7H448V320zm-85.3-85.3H320v42.7h42.7v-42.7zm-256 0H64v42.7h42.7v-42.7zm0 85.3H64v42.7h42.7V320zm320-320v85.3h-85.3V0H170.7v85.3H85.3V0H0v512h512V0h-85.3zm42.6 469.3H42.7V128h426.7v341.3zm-362.6-64H64V448h42.7v-42.7zM277.3 320h-42.7v42.7h42.7V320zM192 405.3h-42.7V448H192v-42.7zm85.3 0h-42.7V448h42.7v-42.7zm-85.3-256h-42.7V192H192v-42.7zm0 85.4h-42.7v42.7H192v-42.7zm0 85.3h-42.7v42.7H192V320zM405.3 0h-42.7v64h42.7V0z"/>`),
		g.Group(children),
	)
}