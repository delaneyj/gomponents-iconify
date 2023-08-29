package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1536 1184V480q0-40-28-68t-68-28H736q-40 0-68-28t-28-68v-64q0-40-28-68t-68-28H224q-40 0-68 28t-28 68v960q0 40 28 68t68 28h1216q40 0 68-28t28-68zm128-704v704q0 92-66 158t-158 66H224q-92 0-158-66T0 1184V224q0-92 66-158T224 0h320q92 0 158 66t66 158v32h672q92 0 158 66t66 158z"/>`),
		g.Group(children),
	)
}