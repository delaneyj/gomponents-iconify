package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FSixKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" rx="3"/><path stroke="#fff" d="M30 32C32.2091 32 34 30.2091 34 28C34 25.7909 32.2091 24 30 24C27.7909 24 26 25.7909 26 28C26 30.2091 27.7909 32 30 32Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M34 18.8763C33.3821 17.1893 31.8922 16 30.1524 16C27.8591 16 26 17.451 26 20V28"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M21 16H14V32"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M14 24H21"/></g>`),
		g.Group(children),
	)
}