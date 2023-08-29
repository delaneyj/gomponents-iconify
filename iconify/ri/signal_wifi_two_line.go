package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalWifiTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.001 3c4.284 0 8.22 1.497 11.31 3.996L12 21L.691 6.997A17.925 17.925 0 0 1 12 3Zm0 9c-1.42 0-2.764.33-3.959.915l3.959 4.902l3.958-4.902A8.963 8.963 0 0 0 12.001 12Zm0-7c-3.028 0-5.923.842-8.42 2.392l3.178 3.935A10.95 10.95 0 0 1 12.001 10c1.898 0 3.683.48 5.242 1.327L20.42 7.39A15.921 15.921 0 0 0 12 5Z"/>`),
		g.Group(children),
	)
}