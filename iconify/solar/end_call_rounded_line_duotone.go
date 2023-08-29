package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EndCallRoundedLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.083 10.503C2.277 11.373 2 12.613 2 13.851c0 2.06 1.782 3.564 3.607 3.046l1.34-.38C8.156 16.175 9 14.983 9 13.62m11.917-3.116c.806.87 1.083 2.11 1.083 3.348c0 2.06-1.782 3.564-3.607 3.046l-1.34-.38C15.844 16.175 15 14.983 15 13.62"/><path d="M3.083 10.503C4.44 9.04 7.193 7 12 7c4.807 0 7.56 2.039 8.917 3.503M9 13.618s0-1.654 3-1.654s3 1.654 3 1.654" opacity=".5"/></g>`),
		g.Group(children),
	)
}