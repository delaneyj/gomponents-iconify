package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Puppet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.984 21.611H6.595v-2.388h2.39zM6.595 2.39h2.39v2.388h-2.39zm13.198 6.028h-5.48l.001-.002l-2.941-2.941V0H4.207v7.166h5.48l2.938 2.938l.002-.001v3.794l-.003-.003l-2.94 2.94H4.207V24h7.166v-5.477l2.94-2.94h5.48V8.417"/>`),
		g.Group(children),
	)
}