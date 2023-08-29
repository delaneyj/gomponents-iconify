package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresentationTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1877 0v171h-170v853q0 35-13 66t-37 55t-55 36t-66 14h-512v512h341v170H512v-170h341v-512H341q-35 0-66-13t-54-37t-36-54t-14-67V171H0V0h1877zm-341 1024V171H341v853h1195zm-171-512v171H512V512h853z"/>`),
		g.Group(children),
	)
}