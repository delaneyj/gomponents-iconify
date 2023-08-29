package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Door(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="m4.083.041l6.959 2.667v12.084l-6.709-1.709l-.208.834l7.803 2v-2.042l1.968-.001V.041H4.083z"/><path d="M9 8h.875v.875H9z"/></g>`),
		g.Group(children),
	)
}