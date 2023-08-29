package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.48 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33.04a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.398 22.95a3.229 3.229 0 0 1 3.228-3.228h0m-3.228 0v8.556m-2.492-3.228a3.229 3.229 0 0 1-3.229 3.228h0a3.229 3.229 0 0 1-3.228-3.229v-2.098a3.229 3.229 0 0 1 3.228-3.23h0a3.229 3.229 0 0 1 3.23 3.23m-.001 5.327v-8.556M9.5 22.95a3.229 3.229 0 0 1 3.229-3.228h0a3.229 3.229 0 0 1 3.228 3.229v2.098a3.229 3.229 0 0 1-3.228 3.23h0a3.229 3.229 0 0 1-3.229-3.23m0 3.229V15.363m29 9.687a3.229 3.229 0 0 1-3.229 3.228h0a3.229 3.229 0 0 1-3.228-3.229v-2.098a3.229 3.229 0 0 1 3.228-3.23h0a3.229 3.229 0 0 1 3.229 3.23m0-3.229v12.915"/>`),
		g.Group(children),
	)
}