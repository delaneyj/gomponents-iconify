package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DivingSuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M17 24H14.09C8.52 24 4 19.5203 4 14C4 8.47968 8.52 4 14.09 4L20 9L25.91 4C31.48 4 36 8.47968 36 14C36 19.5203 31.48 24 25.91 24H23"/><path fill="#2F88FF" d="M26 36H14V44H26V36Z"/><path d="M26 40H36C40.42 40 44 36.42 44 32V14H36"/><path fill="#2F88FF" d="M21.2198 30H19.9998H18.7798C17.3198 30 16.2098 28.69 16.4498 27.26L17.6698 19.97C17.8598 18.83 18.8398 18 19.9998 18C21.1498 18 22.1398 18.83 22.3298 19.97L23.5498 27.26C23.7798 28.69 22.6698 30 21.2198 30Z"/></g>`),
		g.Group(children),
	)
}