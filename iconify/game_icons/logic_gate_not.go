package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogicGateNot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M105 111.3v289.4L365.5 256ZM16 247v18h71v-18zm400-14c-12.8 0-23 10.2-23 23s10.2 23 23 23s23-10.2 23-23s-10.2-23-23-23zm40 14c.6 2.9 1 5.9 1 9c0 3.1-.4 6.1-1 9h40v-18z"/>`),
		g.Group(children),
	)
}