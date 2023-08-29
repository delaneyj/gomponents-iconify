package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maximize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#597B91" d="M463.447 58.145H53.474c-7.858 0-14.228 6.37-14.228 14.228V454.75c0 6.577 5.692 11.91 12.714 11.91h413.002c7.022 0 12.714-5.332 12.714-11.91V72.372c-.001-7.857-6.371-14.227-14.229-14.227z"/><path fill="#E0E0E0" d="M61.234 127.695v307.186c0 6.325 5.122 11.452 11.44 11.452h371.63c6.318 0 11.44-5.127 11.44-11.452V127.695H61.234z"/>`),
		g.Group(children),
	)
}