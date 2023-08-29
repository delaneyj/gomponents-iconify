package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkedinSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 13.5v-12ZM5 5H4V4h1v1Zm-1 6V6h1v5H4Zm4.5-4A1.5 1.5 0 0 0 7 8.5V11H6V6h1v.5a2.5 2.5 0 0 1 4 2V11h-1V8.5A1.5 1.5 0 0 0 8.5 7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}