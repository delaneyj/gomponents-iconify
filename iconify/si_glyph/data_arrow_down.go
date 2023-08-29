package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.997 10.04H2.023c-.555 0-1.003.833-1.003 1.859v1.241c0 1.026.448 1.859 1.003 1.859h13.974c.554 0 1.003-.833 1.003-1.859v-1.241c0-1.026-.449-1.859-1.003-1.859zM12.5 14.104c-.885 0-1.604-.732-1.604-1.635s.719-1.636 1.604-1.636c.885 0 1.604.732 1.604 1.636c0 .902-.719 1.635-1.604 1.635zm-.734-9.94H9.878V1.687c0-.557-.421-.633-.938-.633c-.519 0-.938.076-.938.633v2.477H6.213a.735.735 0 0 0 0 1.033l2.265 2.568a.725.725 0 0 0 1.027 0l2.261-2.568a.73.73 0 0 0 0-1.033z"/>`),
		g.Group(children),
	)
}