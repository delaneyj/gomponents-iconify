package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M1 8.041C1 3.652 4.582.082 8.985.082c4.401 0 7.983 3.57 7.983 7.959c0 4.389-3.582 7.959-7.983 7.959C4.582 16 1 12.43 1 8.041zm14.057 0c0-3.333-2.715-6.045-6.051-6.045c-3.337 0-6.052 2.712-6.052 6.045c0 3.333 2.716 6.045 6.052 6.045c3.337 0 6.051-2.712 6.051-6.045z"/><path d="M8.975 5.02L5.071 8.022l3.905 2.951V8.97h3.14c.345 0 .791-.324.791-.955c0-.63-.483-.975-.826-.975H8.976V5.02h-.001z"/></g>`),
		g.Group(children),
	)
}