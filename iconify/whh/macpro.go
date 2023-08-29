package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Macpro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.488 1024h-64q-16 0-33.5-5.5t-22.5-13.5l-67-109h-522l-67 109q-5 8-23 13.5t-33 5.5h-64q-27 0-45.5-19t-18.5-45V64q0-27 18.5-45.5T64.488 0h64q15 0 33 5.5t23 13.5l67 109h522l67-109q5-8 22.5-13.5t33.5-5.5h64q26 0 45 18.5t19 45.5v896q0 27-19 45.5t-45 18.5zm-448-576q-26 0-45 19t-19 45t19 45t45 19t45-18.5t19-45.5t-19-45.5t-45-18.5z"/>`),
		g.Group(children),
	)
}