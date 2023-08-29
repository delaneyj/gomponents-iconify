package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.893 2.809a.499.499 0 0 0-.393-.808h-11a.501.501 0 0 0-.393.808L7 9.037V15H5.5a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1H9V9.037l4.893-6.228zM12.471 3L10.9 5H5.1L3.529 3h8.943z"/>`),
		g.Group(children),
	)
}