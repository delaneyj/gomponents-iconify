package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 2.5h10a.5.5 0 0 1 .5.5v5.063A2.004 2.004 0 0 0 13 8H3c-.173 0-.34.022-.5.063V3a.5.5 0 0 1 .5-.5ZM2.5 10v3a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-.5-.5H3a.5.5 0 0 0-.5.5ZM1 10V3a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2v-3Zm11 1.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-4 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}