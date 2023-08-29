package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mzabaps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.626 31.396c0-5.001-18.806-5.001-18.806-5.001V43.5h18.806V31.396Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.626 43.5H39.18V22.894l-2.1-10.204l-2.313 10.573c-5.195.71-12.024 2.35-14.532 4.067M16.897 7.933s-1.65-.39-1.578-1.218c.03-.339 3.259-.798 5.62-.743A7.583 7.583 0 0 0 16.448 4.5a7.628 7.628 0 0 0-7.628 7.628c0 1.76.602 3.375 1.603 4.666c1.587-2.08 4.888-6.45 6.474-8.86Zm5.878-.067c-1.958 2.488-8.257 7.188-11.779 9.592a7.6 7.6 0 0 0 5.452 2.298a7.628 7.628 0 0 0 6.327-11.89Zm0 0a8.13 8.13 0 0 0 .587-.856M10.03 18.025c.484-.166.966-.567.966-.567"/>`),
		g.Group(children),
	)
}