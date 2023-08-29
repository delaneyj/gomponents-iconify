package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickyNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1024 1120v416H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h1344q40 0 68 28t28 68v928h-416q-40 0-68 28t-28 68zm128 32h381q-15 82-65 132l-184 184q-50 50-132 65v-381z"/>`),
		g.Group(children),
	)
}