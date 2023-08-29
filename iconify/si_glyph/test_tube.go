package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.921.969c0-.937-.091-.938-1.05-.938H5.13c-.958 0-1.052.001-1.052.938h.953v2.94L2.019 14.238c0 .939.777 1.699 1.736 1.699h9.489c.958 0 1.736-.76 1.736-1.699L11.992 3.879l-.014-2.91h.943zm1.048 12.637c.271.884-.203 1.435-1.432 1.435H4.562c-1.401 0-1.745-.593-1.432-1.435l2.786-9.652L5.905.941h5.125V3.93l2.939 9.676z"/><path d="M10.039 6.031h-3.02l-1.838 6.308c-.355 1.15-.24 1.584 1.408 1.584h3.879c1.633 0 1.67-.518 1.409-1.584l-1.838-6.308z"/></g>`),
		g.Group(children),
	)
}