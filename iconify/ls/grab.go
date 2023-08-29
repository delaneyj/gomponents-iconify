package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M298 14h85v42h-85V14zM128 56h85v43h-85V56zm85 43h42V56h43v213h-43v-85h-42V99zm212-43h85v43h-85v170h-42V56h42zM85 184V99h43v85H85zm468-85v85h42v42h-42v85h-43V99h43zm42 85v-43h43v43h-43zm43 42v-42h42v170h-42V226zm-468 43h-42v-85h42v85zm-127 0h85v42H43v-42zM0 396v-85h43v85H0zm128-42v-43h42v-42h43v127h-43v-42h-42zm467 85v-85h43v127h-43v-42zM85 396v43H43v-43h42zm43 128H85v-85h43v85zm425 42v-85h42v85h-42zm-383-42v42h-42v-42h42zm43 42v43h-43v-43h43zm297 86v-86h43v128h-43v-42zm-255 42h-42v-85h42v85z"/>`),
		g.Group(children),
	)
}