package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplaneDeparture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#9b9b9a" d="M24 39.5L8.119 34.508l-4.437 1.375L13.3 45"/><path fill="#d0cfce" d="m42.589 30.5l13.635-4.637s13.5-.229 10.33 3.255c-3.171 3.484-25.426 12.196-32.002 14.651c-6.577 2.455-14.712 3.75-20.337 3.86c-4.924.095.59-2.892.59-2.892s6.255-3.55 9.024-4.757l2.635-1.949"/><path fill="#9b9b9a" d="m48.2 35l-25.755-2.359l-2.093 1.493L38 40.5"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M19.974 39.139L7.542 34.736l-3.067 1.073l8.169 7.322m34.407-8.438l-24.606-2.052l-2.093 1.493l16.754 5.979"/><path d="m43.174 31.323l11.15-4.486a10.312 10.312 0 0 1 4.172-.74c3.7.113 10.39.667 8.057 3.23c-3.17 3.484-25.425 12.196-32.001 14.651s-14.712 3.75-20.337 3.86c-4.924.094.59-2.892.59-2.892s6.255-3.55 9.024-4.757l2.784-1.213"/></g>`),
		g.Group(children),
	)
}