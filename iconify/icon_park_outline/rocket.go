package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M18.705 7.894L24 4l5.295 3.894A14 14 0 0 1 35 19.172V37H13V19.172a14 14 0 0 1 5.705-11.278Z"/><path stroke-linecap="round" d="m13 17l-6 6v8l6-3V17Zm22 0l6 6v8l-6-3V17Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M18 39v3m6-3v5m6-5v3"/></g>`),
		g.Group(children),
	)
}