package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapLocationSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.75 9a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M11.838 2.5a6.153 6.153 0 0 0-6.132 5.648A6.645 6.645 0 0 0 7.184 12.9l3.595 4.396a1.578 1.578 0 0 0 2.442 0l3.595-4.396a6.644 6.644 0 0 0 1.478-4.752A6.153 6.153 0 0 0 12.162 2.5h-.324ZM12 5.25a3.75 3.75 0 1 0 0 7.5a3.75 3.75 0 0 0 0-7.5Z" clip-rule="evenodd"/><path fill="currentColor" d="M7.335 15.33a.75.75 0 0 1 .336 1.005L6.214 19.25h11.572l-1.457-2.915a.75.75 0 1 1 1.342-.67l2 4A.75.75 0 0 1 19 20.75H5a.75.75 0 0 1-.67-1.085l2-4a.75.75 0 0 1 1.005-.336Z"/>`),
		g.Group(children),
	)
}