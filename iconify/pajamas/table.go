package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 4.5v-2h11v2h-11Zm0 1.5v3h4.75V6H2.5Zm6.25 0v3h4.75V6H8.75ZM2.5 10.5h4.75v3H2.5v-3Zm6.25 0v3h4.75v-3H8.75ZM1 2a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}