package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M.076 7.989c0 4.401 3.562 7.969 7.955 7.969c4.394 0 7.955-3.567 7.955-7.969S12.424.02 8.031.02C3.639.021.076 3.588.076 7.989zm3.914 2.105c0-2.528 1.169-4.264 3.785-4.264l.256.001V4.198c.213-.214.463-.179.678.036l3.229 2.422a.552.552 0 0 1 0 .775L8.709 9.853a.488.488 0 0 1-.678-.035V8.22l-.236-.001c-1.816 0-2.439.677-3.232 2.126c-.167.308-.573.63-.573-.251z"/>`),
		g.Group(children),
	)
}