package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneLandscapeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<rect width="256" height="480" x="128" y="16" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" rx="48" ry="48" transform="rotate(-90 256 256)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M16 336v-24a8 8 0 0 1 8-8h0a16 16 0 0 0 16-16v-64a16 16 0 0 0-16-16h0a8 8 0 0 1-8-8v-24"/>`),
		g.Group(children),
	)
}