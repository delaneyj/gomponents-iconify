package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShuffleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M40 33L44 37L40 41"/><path stroke-linejoin="round" d="M40 7L44 11L40 15"/><path d="M44 11H37C29.8203 11 24 16.8203 24 24C24 31.1797 29.8203 37 37 37H44"/><path d="M4 37H11C18.1797 37 24 31.1797 24 24C24 16.8203 18.1797 11 11 11H4"/></g>`),
		g.Group(children),
	)
}