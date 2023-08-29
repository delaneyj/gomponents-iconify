package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chiki(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.62 34.868s-3.04-4.812-1.838-9.223s8.532-6.705 14.257-3.152c5.67 3.553 5.803 10.003-.344 10.08c-6.639.082-9.791.143-12.03 2.295Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.777 20.957c.294-1.145 3.27-11.82 11.894-11.82c9.111 0 14.369 6.45 16.485 7.876a12.43 12.43 0 0 0 4.344 2.072a7.256 7.256 0 0 1-5.725 3.998"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.594 19.274s1.893 4.979 1.259 9.056c-.63 4.065 2.573 7.106 2.573 7.106s-3.955.516-4.868-.515m1.036 2.286a10.268 10.268 0 0 0-7.73.516c-3.509 1.927-12.921 2.05-17.042-4.79"/><ellipse cx="18.724" cy="19.999" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.873" ry=".762" transform="rotate(-80.65 18.724 20)"/><ellipse cx="27.051" cy="20.602" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.147" ry=".874" transform="rotate(-75 27.05 20.603)"/>`),
		g.Group(children),
	)
}