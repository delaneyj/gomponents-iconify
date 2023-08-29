package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M274.6 25.623a32.006 32.006 0 0 0-37.2 0L16 183.766V496h480V183.766ZM464 402.693L339.97 322.96L464 226.492ZM256 51.662L454.429 193.4L311.434 304.615L256 268.979l-55.434 35.636L57.571 193.4ZM48 226.492l124.03 96.468L48 402.693ZM464 464H48v-23.265l208-133.714l208 133.714Z"/>`),
		g.Group(children),
	)
}