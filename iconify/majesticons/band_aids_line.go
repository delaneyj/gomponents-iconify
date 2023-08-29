package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BandAidsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="9" height="19" x="2" y="8.364" rx="4.5" transform="rotate(-45 2 8.364)"/><path d="m11.9 18.264l-.354.353a4.5 4.5 0 0 1-6.364 0v0a4.5 4.5 0 0 1 0-6.364l.354-.353M11.9 5.536l.353-.354a4.5 4.5 0 0 1 6.364 0v0a4.5 4.5 0 0 1 0 6.364l-.354.354m-8.484 0h0M11.9 9.778h0m2.121 2.122h0M11.9 14.021h0"/></g>`),
		g.Group(children),
	)
}