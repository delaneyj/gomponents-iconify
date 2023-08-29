package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" d="M13 42C22.9411 42 31 33.9411 31 24C31 14.0589 22.9411 6 13 6"/><circle cx="9" cy="24" r="3" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 24L42 24"/><path stroke-linecap="round" stroke-linejoin="round" d="M38 20L42 24L38 28"/></g>`),
		g.Group(children),
	)
}