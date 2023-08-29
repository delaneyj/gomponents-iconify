package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeLightroomSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.77 3.082a47.472 47.472 0 0 1 10.46 0c1.899.212 3.43 1.707 3.653 3.613a45.67 45.67 0 0 1 0 10.61c-.223 1.906-1.754 3.401-3.652 3.614a47.468 47.468 0 0 1-10.461 0c-1.899-.213-3.43-1.708-3.653-3.613a45.672 45.672 0 0 1 0-10.611C3.34 4.789 4.871 3.294 6.77 3.082ZM12.728 15.5h-4.4V8.086h1.353v6.281h3.047V15.5Zm1.883 0h-1.353v-3.63c0-.726-.011-1.243-.044-1.727h1.177l.044 1.023h.044c.264-.759.89-1.144 1.463-1.144c.132 0 .209.011.319.033v1.276a1.881 1.881 0 0 0-.396-.044c-.65 0-1.09.418-1.21 1.023a2.322 2.322 0 0 0-.044.418V15.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}