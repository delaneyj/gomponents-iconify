package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deepstash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.239 8.01h25.75v25.751"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.75 5.5H42.5v25.75"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 16.75h25.75V42.5H5.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.831 13.418h25.751v25.75"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.704 10.545h25.75v25.751M42.5 31.25L31.25 42.5m-14.5-37L5.5 16.75"/>`),
		g.Group(children),
	)
}