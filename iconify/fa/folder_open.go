package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1879 824q0 31-31 66l-336 396q-43 51-120.5 86.5T1248 1408H160q-34 0-60.5-13T73 1352q0-31 31-66l336-396q43-51 120.5-86.5T704 768h1088q34 0 60.5 13t26.5 43zm-343-344v160H704q-94 0-197 47.5T343 807L6 1203l-5 6q0-4-.5-12.5T0 1184V224q0-92 66-158T224 0h320q92 0 158 66t66 158v32h544q92 0 158 66t66 158z"/>`),
		g.Group(children),
	)
}