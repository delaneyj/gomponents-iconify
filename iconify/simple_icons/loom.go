package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 10.665h-7.018l6.078-3.509l-1.335-2.312l-6.078 3.509l3.508-6.077L16.843.94l-3.508 6.077V0h-2.67v7.018L7.156.94L4.844 2.275l3.509 6.077l-6.078-3.508L.94 7.156l6.078 3.509H0v2.67h7.017L.94 16.844l1.335 2.313l6.077-3.508l-3.509 6.077l2.312 1.335l3.509-6.078V24h2.67v-7.017l3.508 6.077l2.312-1.335l-3.509-6.078l6.078 3.509l1.335-2.313l-6.077-3.508h7.017v-2.67H24zm-12 4.966a3.645 3.645 0 1 1 0-7.29a3.645 3.645 0 0 1 0 7.29z"/>`),
		g.Group(children),
	)
}