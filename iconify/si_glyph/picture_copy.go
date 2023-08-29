package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureCopy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M14.969.031H5v2.938h1V1h8.027v7.001H11v2.951h3.969V.031z"/><path d="M0 4v11h10V4H0zm8.967 8h-8V5h8v7z"/></g>`),
		g.Group(children),
	)
}