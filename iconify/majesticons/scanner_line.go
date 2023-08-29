package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScannerLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 17l-2-7H5l-2 7m18 0H3m18 0l-.621 2.485A2 2 0 0 1 18.439 21H5.561a2 2 0 0 1-1.94-1.515L3 17m5.309-3h7.382a.5.5 0 0 0 .447-.724v0a.5.5 0 0 0-.447-.276H8.309a.5.5 0 0 0-.447.276v0a.5.5 0 0 0 .447.724zM3.728 5.55A2 2 0 0 1 5.651 3H18.35a2 2 0 0 1 1.923 2.55L19 10H5L3.728 5.55z"/>`),
		g.Group(children),
	)
}