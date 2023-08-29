package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M52 2H12C6.479 2 2 6.477 2 12v40c0 5.523 4.479 10 10 10h40c5.523 0 10-4.477 10-10V12c0-5.523-4.477-10-10-10zm5 43.666A8.332 8.332 0 0 1 48.668 54H15.334A8.334 8.334 0 0 1 7 45.666V12.334A8.334 8.334 0 0 1 15.334 4h33.334A8.332 8.332 0 0 1 57 12.334v33.332z"/><path fill="currentColor" d="M33.662 45v-6.414H20v-5.348L34.482 13h5.373v20.217H44v5.369h-4.145V45h-6.193zm0-11.783V22.326l-7.674 10.891h7.674z"/>`),
		g.Group(children),
	)
}