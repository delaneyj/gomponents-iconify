package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func COutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18.26 10.142a15 15 0 0 1 16.347 3.251a3 3 0 1 1-4.243 4.243a9.001 9.001 0 1 0 0 12.728a3 3 0 1 1 4.243 4.243A15 15 0 1 1 18.26 10.142Zm8.276 1.108a13 13 0 1 0 6.656 21.942a1 1 0 0 0-1.414-1.414a11 11 0 1 1 0-15.556a1 1 0 1 0 1.414-1.414a13 13 0 0 0-6.656-3.558Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}