package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cutlery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13 .8c0-.5-.4-.8-.8-.8H12c-1.7 0-3 1.9-3 4.7v.9c0 1 .5 1.9 1.4 2.4c-.3 1.2-.4 2.5-.4 2.5v4c0 .8.7 1.5 1.5 1.5s1.5-.7 1.5-1.5v-4c0-.4-.1-1.4-.3-2.3c.2-.2.3-.4.3-.7V.8zM7.2 0H7v3.5c0 .3-.2.5-.5.5S6 3.8 6 3.5v-3c0-.3-.2-.5-.5-.5S5 .2 5 .5v3c0 .3-.2.5-.5.5S4 3.8 4 3.5V0h-.2c-.4 0-.8.4-.8.8v3.7c0 1 .6 1.9 1.5 2.3c-.4 1.6-.5 3.7-.5 3.7v4c0 .8.7 1.5 1.5 1.5S7 15.3 7 14.5v-4c0-.5-.1-2.3-.4-3.7C7.4 6.4 8 5.5 8 4.5V.8c0-.4-.4-.8-.8-.8z"/>`),
		g.Group(children),
	)
}