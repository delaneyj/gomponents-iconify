package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 11.5a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9ZM7 13c1.387 0 2.663-.47 3.68-1.26l3.04 3.04a.75.75 0 1 0 1.06-1.06l-3.04-3.04A6 6 0 1 0 7 13Zm0-3a.75.75 0 0 1-.75-.75v-1.5h-1.5a.75.75 0 0 1 0-1.5h1.5v-1.5a.75.75 0 0 1 1.5 0v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5A.75.75 0 0 1 7 10Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}