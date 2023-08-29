package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 7a6 6 0 0 1 6-6h2a6 6 0 0 1 6 6v6a6.002 6.002 0 0 1-6 6v1h5v2H6v-2h5v-1a6.002 6.002 0 0 1-6-6V7Zm2 4.5V13h3v2H7.535A4 4 0 0 0 11 17h2a4 4 0 0 0 3.465-2H14v-2h3v-1.5h-3v-2h3V8h-3V6h2.874A4.002 4.002 0 0 0 13 3h-2a4.002 4.002 0 0 0-3.874 3H10v2H7v1.5h3v2H7Z"/>`),
		g.Group(children),
	)
}