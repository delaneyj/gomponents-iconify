package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v11m18-4v7a2 2 0 0 1-2 2h-3m5-9c-6.442 0-10.105 1.985-12.055 4.243M3 16v3a2 2 0 0 0 2 2v0h11M3 16c1.403-.234 3.637-.293 5.945.243M16 21c-1.704-2.768-4.427-4.148-7.055-4.757M8.5 7C8 7 7 7.3 7 8.5S8 10 8.5 10S10 9.7 10 8.5S9 7 8.5 7z"/>`),
		g.Group(children),
	)
}