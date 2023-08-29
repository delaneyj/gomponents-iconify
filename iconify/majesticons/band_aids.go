package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BandAids(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><rect width="9" height="19" x="2" y="8.364" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="4.5" transform="rotate(-45 2 8.364)"/><path fill="currentColor" d="m11.546 18.617l.354-.353L5.535 11.9l-.354.353a4.5 4.5 0 0 0 6.364 6.364zm.707-13.435l-.354.354l6.365 6.364l.353-.354a4.5 4.5 0 1 0-6.364-6.364z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11.546 18.617l.354-.353L5.535 11.9l-.354.353a4.5 4.5 0 0 0 6.364 6.364zm.707-13.435l-.354.354l6.365 6.364l.353-.354a4.5 4.5 0 1 0-6.364-6.364zM9.779 11.9h0M11.9 9.778h0m2.121 2.122h0M11.9 14.021h0"/></g>`),
		g.Group(children),
	)
}