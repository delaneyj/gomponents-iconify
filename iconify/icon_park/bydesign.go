package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bydesign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8.00002 7.28206L19.5556 4L19.5555 16.7179L8 20L8.00002 7.28206Z"/><path d="M8.00002 25.6752L19.5556 22.3932L19.5555 35.1111L8 38.3932L8.00002 25.6752Z"/><path d="M24.8887 21.0759L36.4443 18.2222L36.4442 41.1462L24.8887 43.9999L24.8887 21.0759Z"/></g>`),
		g.Group(children),
	)
}