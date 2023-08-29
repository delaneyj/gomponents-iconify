package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M302.2 0H67.5v512h384V149.3H302.2V0zm64 320v42.7h-85.3V448h-42.7v-85.3h-85.3V320h85.3v-85.3h42.7V320h85.3zM323.5 0v128h128L323.5 0z"/>`),
		g.Group(children),
	)
}