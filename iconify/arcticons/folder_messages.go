package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderMessages(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.5a3 3 0 0 1 3-3h8.718a4 4 0 0 1 2.325.745l4.914 3.51a4 4 0 0 0 2.325.745H40.5a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3h-33a3 3 0 0 1-3-3v-25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16 20.853v11.292c0 .95 1.149 1.426 1.82.754l2.447-2.446h10.666c.59 0 1.067-.478 1.067-1.067v-8.533c0-.589-.478-1.067-1.067-1.067H17.067c-.59 0-1.067.478-1.067 1.067Z"/>`),
		g.Group(children),
	)
}