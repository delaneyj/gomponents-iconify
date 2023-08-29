package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpenO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1781 803q0-35-53-35H640q-40 0-85.5 21.5T483 842l-294 363q-18 24-18 40q0 35 53 35h1088q40 0 86-22t71-53l294-363q18-22 18-39zM640 640h768V480q0-40-28-68t-68-28H736q-40 0-68-28t-28-68v-64q0-40-28-68t-68-28H224q-40 0-68 28t-28 68v853l256-315q44-53 116-87.5T640 640zm1269 163q0 62-46 120l-295 363q-43 53-116 87.5t-140 34.5H224q-92 0-158-66T0 1184V224q0-92 66-158T224 0h320q92 0 158 66t66 158v32h544q92 0 158 66t66 158v160h192q54 0 99 24.5t67 70.5q15 32 15 68z"/>`),
		g.Group(children),
	)
}