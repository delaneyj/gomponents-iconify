package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M217 0c116 0 209 95 209 209c0 115-213 559-213 559S0 324 0 209C0 95 93 0 209 0h8zm-4 287c52 0 95-44 95-97c0-52-43-95-95-95c-53 0-96 43-96 95c0 53 43 97 96 97z"/>`),
		g.Group(children),
	)
}