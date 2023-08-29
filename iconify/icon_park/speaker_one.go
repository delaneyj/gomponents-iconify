package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="11" height="18" x="4" y="15" fill="#2F88FF"/><path d="M15 15L33 8V40L15 33"/><path d="M40 17H42"/><path d="M39 25H44"/><path d="M40 33H42"/></g>`),
		g.Group(children),
	)
}