package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HostSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<g id="clarityHostSolidBadged0" fill="currentColor"><circle cx="18" cy="25.64" r="1.5"/><path d="M24 10.49v1.17H12v-1.6h11.7a7.42 7.42 0 0 1-1-2.4H12v-1.6h10.5V6a7.45 7.45 0 0 1 1.25-4.14H9.5A1.5 1.5 0 0 0 8 3.36v30.5h20V13.22a7.5 7.5 0 0 1-4-2.73Zm-6 18.15a3 3 0 1 1 3-3a3 3 0 0 1-3 3Z"/><circle cx="30" cy="6" r="5"/></g>`),
		g.Group(children),
	)
}