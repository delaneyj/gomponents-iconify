package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersMinimalisticBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M22 12s-.993.89-2.979 1.685l-2.808 1.123C14.227 15.604 13.233 16 12 16c-1.234 0-2.227-.397-4.213-1.191L4.98 13.684C2.993 12.891 2 12 2 12m0 4s.993.89 2.979 1.685l2.808 1.123C9.773 19.604 10.767 20 12 20c.954 0 1.764-.237 3-.712m4.021-1.603C21.007 16.891 22 16 22 16M4.979 6.315C2.993 7.109 2 7.506 2 8s.993.89 2.979 1.685l2.808 1.123C9.773 11.604 10.767 12 12 12c1.234 0 2.227-.397 4.213-1.191l2.808-1.124C21.007 8.891 22 8.494 22 8s-.993-.89-2.979-1.685l-2.808-1.123C14.227 4.397 13.233 4 12 4c-.954 0-1.764.237-3 .712"/>`),
		g.Group(children),
	)
}