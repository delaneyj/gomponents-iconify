package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.021 4a1.97 1.97 0 0 0-1.706 1H9.708a1.985 1.985 0 0 0-3.697 1c0 .009.002.018.003.026l-2.052 1.23a1.97 1.97 0 0 0-.961-.257a2 2 0 1 0 2 2c0-.011-.003-.021-.003-.032l2.06-1.235c.281.152.6.247.942.247c.731 0 1.362-.396 1.708-.979h3.607c.344.596.976 1 1.706 1A1.99 1.99 0 0 0 17 6c0-1.105-.887-2-1.979-2z"/>`),
		g.Group(children),
	)
}