package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickyNoteO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1400 1152h-248v248q29-10 41-22l185-185q12-12 22-41zm-280-128h288V128H128v1280h896v-288q0-40 28-68t68-28zm416-928v1024q0 40-20 88t-48 76l-184 184q-28 28-76 48t-88 20H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h1344q40 0 68 28t28 68z"/>`),
		g.Group(children),
	)
}