package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mupdf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.45 5.5a2 2 0 0 0-1.95 2v33.1a2 2 0 0 0 2 2h13v-6.02a2 2 0 0 1 2-1.95h3.16a2 2 0 0 1 2 1.95v5.92h13a2 2 0 0 0 2-2V7.45a2 2 0 0 0-2-1.95ZM24 10.76a10.35 10.35 0 1 1-10.35 10.35A10.35 10.35 0 0 1 24 10.76Zm0 5.39a5 5 0 0 0 0 9.92h0a5 5 0 0 0 5-5h0a5 5 0 0 0-5-4.92Z"/>`),
		g.Group(children),
	)
}