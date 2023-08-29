package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontSizeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M585 512h143l582 1536h-158l-188-512H343l-184 512H0L585 512zm-198 897h529L683 781q-20-53-28-108h-5q-11 56-30 108l-233 628zM1604 283l-230 230l-90-91l320-320l320 320l-91 91l-229-230zm-4 587l230-230l90 90l-320 320l-320-320l90-90l230 230z"/>`),
		g.Group(children),
	)
}