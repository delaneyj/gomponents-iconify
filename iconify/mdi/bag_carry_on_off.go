package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BagCarryOnOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.1 4.9l3.9 4V19c0 .5.2 1 .6 1.4c.4.4.9.6 1.4.6V10.8l1 1V21h4.2c-.1-.4-.2-.8-.2-1c0-1.2.5-2 1.6-2.6l.8.8c-.9.3-1.3 1-1.3 1.9c0 .5.2 1 .6 1.4c.3.3.8.5 1.3.5c.9 0 1.6-.4 1.9-1.3l1.2 1.2l1.4-1.4l-17-17l-1.4 1.4M12 2c0 .5.2 1 .6 1.4s.9.6 1.4.6v3H9.8l6.2 6.2V2h-4Z"/>`),
		g.Group(children),
	)
}