package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pointer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M298 14v42h-85V14h85zm0 170V56h42v128h170v42h-42v85h-43v-85h-85v85h-42V184zm-170 85v42H43v-42h85zm425 42h-43v-85h85v43h-42v42zm42 0v-42h43v212h-43V311zM0 311h43v85H0v-85zm170 85v-42h-42v-43h42V56h43v340h-43zM43 439v-43h42v43H43zm255 42v-85h42v170h-42v-85zm85 0v-85h42v170h-42v-85zM85 439h43v85H85v-85zm425 0v127h-42V396h42v43zm43 127v-85h42v85h-42zm-425 0v-42h42v42h-42zm42 43v-43h43v43h-43zm340 43v-86h43v128h-85v-42h42zm-85 0v-43h43v43h-43zm-127 0h127v42H213v-85h42v43h43z"/>`),
		g.Group(children),
	)
}