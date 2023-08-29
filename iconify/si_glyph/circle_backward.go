package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.031.021c-4.394 0-7.955 3.567-7.955 7.969s3.562 7.969 7.955 7.969c4.393 0 7.955-3.567 7.955-7.969S13.424.021 9.031.021zM12.5 10.346c-.793-1.449-1.416-2.126-3.232-2.126l-.236.001v1.598a.488.488 0 0 1-.678.035L5.125 7.432a.552.552 0 0 1 0-.775l3.229-2.422c.215-.215.465-.25.678-.036v1.633l.256-.001c2.616 0 3.785 1.735 3.785 4.264c-.001.88-.406.558-.573.251z"/>`),
		g.Group(children),
	)
}