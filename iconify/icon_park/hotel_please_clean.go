package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotelPleaseClean(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" d="M13 27C13 27 33 27 33 15V44H13V27Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M33 15V14C33 8.47715 28.5228 4 23 4C17.4772 4 13 8.47715 13 14M33 15C33 27 13 27 13 27V44H33V15Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 34L22 37L27 32"/></g>`),
		g.Group(children),
	)
}