package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Turnleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M995 581L581 995q-29 29-69 29t-69-29L29 581Q0 552 0 512t29-69L443 29q29-29 69-29t69 29l414 414q29 29 29 69t-29 69zm-291-5q0-53-37.5-90.5T576 448h-64V335q-22-25-37-8L328 457q-8 10-8 24t8 23l147 130q14 17 37-8V512h64q26 0 45 19t19 45v64h64v-64z"/>`),
		g.Group(children),
	)
}