package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteSquareButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#d0d2d3" d="M63.94 57.047a6.892 6.892 0 0 1-6.892 6.898H6.89A6.895 6.895 0 0 1 0 57.047V6.894A6.892 6.892 0 0 1 6.89 0h50.159a6.892 6.892 0 0 1 6.892 6.894v50.153z"/><path fill="#25333a" d="M46.49 13.459H17.45a3.991 3.991 0 0 0-3.991 3.991v29.039a3.995 3.995 0 0 0 3.991 3.997h29.04a3.995 3.995 0 0 0 3.993-3.997v-29.04a3.992 3.992 0 0 0-3.993-3.991"/>`),
		g.Group(children),
	)
}