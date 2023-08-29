package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Seat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" d="M17 21.458c-4.057 1.274-7 5.065-7 9.542c0 5.523 4.477 10 10 10a9.985 9.985 0 0 0 8.662-5"/><path stroke-linecap="round" stroke-linejoin="round" d="M38 20c-4 0-7.5-.5-14-3v12h14v14"/><circle cx="24" cy="8" r="4"/></g>`),
		g.Group(children),
	)
}