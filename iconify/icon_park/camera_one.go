package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="19" r="14"/><circle cx="24" cy="19" r="6" fill="#2F88FF"/><path d="M17 31L11 43H37L31 31"/></g>`),
		g.Group(children),
	)
}