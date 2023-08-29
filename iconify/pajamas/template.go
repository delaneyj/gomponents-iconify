package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Template(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 2.5h-11V5h11V2.5Zm0 4h-7v7h7v-7ZM5 6.5H2.5v7H5v-7ZM2.5 1H1v14h14V1H2.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}