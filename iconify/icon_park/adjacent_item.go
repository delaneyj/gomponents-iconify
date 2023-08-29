package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjacentItem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" d="M14 29H42V41H14V35V29Z"/><path fill="#2F88FF" d="M14 7H42V19H14V13V7Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 13V19H42V7H14V13ZM14 13H6V35H14M14 35V41H42V29H14V35Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 13H6V35H14"/><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 29H42V41H14V35V29Z"/><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 7H42V19H14V13V7Z"/></g>`),
		g.Group(children),
	)
}