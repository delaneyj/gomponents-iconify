package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BootVolume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M32 22v-2h-3v-2h-2v2h-1v2h1v6c0 1.103.897 2 2 2h3v-2h-3v-6h3zm-9 8h-3c-1.103 0-2-.897-2-2v-6c0-1.103.897-2 2-2h3c1.103 0 2 .897 2 2v6c0 1.103-.897 2-2 2zm-3-8v6h3v-6h-3zm-6 8h-3c-1.103 0-2-.897-2-2v-6c0-1.103.897-2 2-2h3c1.103 0 2 .897 2 2v6c0 1.103-.897 2-2 2zm-3-8v6h3v-6h-3zm-6-2H2v-4H0v14h5c1.103 0 2-.897 2-2v-6c0-1.103-.897-2-2-2zm-3 8v-6h3v6H2zm2-14V5h7.586l4 4H28v7h2V9a2.002 2.002 0 0 0-2-2H16.414L13 3.586A1.987 1.987 0 0 0 11.586 3H4a2.002 2.002 0 0 0-2 2v9h2z"/>`),
		g.Group(children),
	)
}