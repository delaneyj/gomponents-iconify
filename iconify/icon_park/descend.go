package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Descend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M11.549 16.6253L13.0129 11.161C13.0129 11.161 8.21885 12.3128 5.91213 14.0184C3.6054 15.7239 3.25986 19.4167 5.98653 20.9909C8.7132 22.5651 44.1733 39.9362 44.1733 39.9362L41.4054 31.142L11.549 16.6253Z"/><path d="M20 35L26 38"/><path d="M29 25L26 9L22 7L19 20"/></g>`),
		g.Group(children),
	)
}