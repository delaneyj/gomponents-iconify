package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Panzerfaust(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m227.752 139.172l51.998 54.29l13-12.45l-52-54.29l-12.998 12.45zM187.74 256.424l9.683 10.11l56.33-53.952l-9.684-10.11l-56.33 53.952zM385.14 72.18l-38.345 72.733l24.227 25.294l74.315-35.176l-60.198-62.85zm15.994-9.322l54.207 56.595l7.832-35.195l-26.54-27.708l-35.498 6.308zM48.828 433.784l20.75 21.666l9.39-8.992l-20.752-21.666l-9.388 8.992zm22.388-21.442l20.75 21.665l264.318-253.164l-20.75-21.666L71.215 412.342z"/>`),
		g.Group(children),
	)
}