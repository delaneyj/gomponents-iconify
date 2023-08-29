package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M30 12V4H18V12"/><path fill="#2F88FF" stroke="#000" d="M13 12.3735C13 12.1672 13.1672 12 13.3735 12H34.6265C34.8328 12 35 12.1672 35 12.3735V33C35 39.0751 30.0751 44 24 44C17.9249 44 13 39.0751 13 33V12.3735Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M20 21L28 21"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M20 29L28 29"/></g>`),
		g.Group(children),
	)
}