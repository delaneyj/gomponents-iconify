package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Equalizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 33.67h13v4.83h-13zm0-4.84h13v4.83h-13zm0-4.83h13v4.83h-13zm13 9.67h13v4.83h-13zm0-4.84h13v4.83h-13zm0-4.83h13v4.83h-13zm0-4.83h13V24h-13zm13 0h13V24h-13zm-13-4.84h13v4.83h-13zm0-4.83h13v4.83h-13zm13 24.17h13v4.83h-13zm0-4.84h13v4.83h-13zm0-4.83h13v4.83h-13z"/>`),
		g.Group(children),
	)
}