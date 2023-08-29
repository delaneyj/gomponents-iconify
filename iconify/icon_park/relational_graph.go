package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RelationalGraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M4 35H44"/><path d="M24 32C24 27.5817 20.1944 24 15.5 24C10.8056 24 7 27.5817 7 32"/><path stroke-linecap="round" stroke-linejoin="round" d="M41 32C41 20.9543 33.3888 12 24 12C14.6112 12 7 20.9543 7 32"/><circle cx="41" cy="35" r="3" fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="35" r="3" fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round"/><circle cx="7" cy="35" r="3" fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round"/></g>`),
		g.Group(children),
	)
}