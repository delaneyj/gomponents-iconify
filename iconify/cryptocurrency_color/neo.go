package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Neo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#58BF00"/><path fill="#FFF" fill-rule="nonzero" d="m25 22.58l-6.99-3.258v-7.22L25 9.623V22.58zM14.823 26L8 22.821V9.958l6.823 3.18V26zm10.01-16.843l-.113.04l-6.71 2.381l-.168.06l-2.843 1.008l-6.73-3.136l9.573-3.396l.084-.03l.177-.063l.062-.021l6.73 3.136l-.063.021z"/></g>`),
		g.Group(children),
	)
}