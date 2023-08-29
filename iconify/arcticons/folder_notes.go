package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderNotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.5a3 3 0 0 1 3-3h8.718a4 4 0 0 1 2.325.745l4.914 3.51a4 4 0 0 0 2.325.745H40.5a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3h-33a3 3 0 0 1-3-3v-25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.328 19.515c.427-.02.842.15 1.132.464l1.084 1.1a1.51 1.51 0 0 1 0 2.11l-2.02 2.02l-3.21-3.225l2.02-2.02a1.56 1.56 0 0 1 .994-.464v.015Zm.196 5.679L21.35 32.39l-4.324 1.11l1.119-4.323l7.17-7.193"/>`),
		g.Group(children),
	)
}