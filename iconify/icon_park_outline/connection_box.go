package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectionBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M43 4H5c-.552 0-1 .488-1 1.09v9.82c0 .602.448 1.09 1 1.09h38c.552 0 1-.488 1-1.09V5.09C44 4.489 43.552 4 43 4Zm0 28H5c-.552 0-1 .488-1 1.09v9.82c0 .602.448 1.09 1 1.09h38c.552 0 1-.488 1-1.09v-9.82c0-.602-.448-1.09-1-1.09Z"/><path stroke-linecap="round" d="M14 16v8.008l20 .01V32m-16 6h12M18 10h12"/></g>`),
		g.Group(children),
	)
}