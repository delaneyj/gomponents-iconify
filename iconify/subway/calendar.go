package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M362.7 320H320v42.7h42.7V320zm0-85.3H320v42.7h42.7v-42.7zm0-85.4H320V192h42.7v-42.7zm-85.4 85.4h-42.7v42.7h42.7v-42.7zm0-85.4h-42.7V192h42.7v-42.7zm0 170.7h-42.7v42.7h42.7V320zM448 149.3h-42.7V192H448v-42.7zm0 85.4h-42.7v42.7H448v-42.7zm0 85.3h-42.7v42.7H448V320zm-341.3 85.3H64V448h42.7v-42.7zm0-85.3H64v42.7h42.7V320zm170.6 85.3h-42.7V448h42.7v-42.7zm-85.3 0h-42.7V448H192v-42.7zM0 0v512h512V0H0zm469.3 469.3H42.7V128h426.7v341.3zM106.7 234.7H64v42.7h42.7v-42.7zM192 320h-42.7v42.7H192V320zm0-170.7h-42.7V192H192v-42.7zm0 85.4h-42.7v42.7H192v-42.7zm170.7 170.6H320V448h42.7v-42.7z"/>`),
		g.Group(children),
	)
}