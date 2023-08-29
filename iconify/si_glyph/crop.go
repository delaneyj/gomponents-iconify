package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M1 3v2.036h1.978V3H1zm12 11v2.031h1.984V14H13zm-.01-9.015v4.976h1.979V3.039H7.047v1.946h5.942z"/><path d="M6.007 11.041V0H4v12.959h13v-1.918H6.007z"/></g>`),
		g.Group(children),
	)
}