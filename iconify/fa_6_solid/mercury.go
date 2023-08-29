package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mercury(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M72.1 7C85.8-4 106-1.8 117 12c17.6 22 44.7 36 75 36s57.3-14 75-36c11.1-13.8 31.2-16 45-5s16 31.2 5 45c-7.8 9.7-16.6 18.4-26.4 26.1C337.3 109.7 368 163.3 368 224c0 89.1-66.2 162.7-152 174.4V424h32c13.3 0 24 10.7 24 24s-10.7 24-24 24h-32v16c0 13.3-10.7 24-24 24s-24-10.7-24-24v-16h-32c-13.3 0-24-10.7-24-24s10.7-24 24-24h32v-25.6C82.2 386.7 16 313.1 16 224c0-60.7 30.7-114.3 77.5-145.9c-9.8-7.6-18.6-16.4-26.4-26.1c-11.1-13.8-8.8-33.9 5-45zM80 224a112 112 0 1 0 224 0a112 112 0 1 0-224 0z"/>`),
		g.Group(children),
	)
}