package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1664 480v704q0 92-66 158t-158 66H224q-92 0-158-66T0 1184V224q0-92 66-158T224 0h320q92 0 158 66t66 158v32h672q92 0 158 66t66 158z"/>`),
		g.Group(children),
	)
}