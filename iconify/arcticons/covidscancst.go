package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Covidscancst(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.027 15H15.973a.973.973 0 0 0-.973.973v16.054a.973.973 0 0 0 .973.973h16.054a.973.973 0 0 0 .973-.973V15.973a.973.973 0 0 0-.973-.973ZM8.5 24h31m-25-18.5h-7a2 2 0 0 0-2 2v7m37 0v-7a2 2 0 0 0-2-2h-7m0 37h7a2 2 0 0 0 2-2v-7m-37 0v7a2 2 0 0 0 2 2h7"/>`),
		g.Group(children),
	)
}