package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossHair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.5 1C4.364 1 1 4.361 1 8.492s3.364 7.492 7.5 7.492c4.137 0 7.5-3.361 7.5-7.492S12.637 1 8.5 1zm.469 13.945v-1.898h-.938v1.898A6.462 6.462 0 0 1 2.056 9h1.912V8H2.056a6.461 6.461 0 0 1 5.943-5.943v1.928h.969v-1.93A6.462 6.462 0 0 1 14.943 8h-1.928v.969h1.93a6.46 6.46 0 0 1-5.976 5.976z"/><path d="M8.984 6.021h-.941v2.021H6.031v.942h2.012v1.975h.941V8.984h1.954v-.942H8.984V6.021z"/></g>`),
		g.Group(children),
	)
}