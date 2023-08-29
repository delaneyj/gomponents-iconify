package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiskAntenna(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.916 8.877c-2.558 2.559-6.535 2.734-8.877.391l9.269-9.27c2.342 2.344 2.169 6.318-.392 8.879zM4.625 5.958c-.846-.845-.783-2.274.137-3.194c.92-.92 2.35-.98 3.193-.137l-3.33 3.331z"/><path d="M13.644 15.997H6.378a.363.363 0 0 1-.318-.542l3.69-6.582a.366.366 0 0 1 .319-.186h.002a.363.363 0 0 1 .318.189l3.574 6.582a.364.364 0 0 1-.319.539z"/></g>`),
		g.Group(children),
	)
}