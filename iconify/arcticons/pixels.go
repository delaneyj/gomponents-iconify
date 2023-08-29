package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixels(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.91 33.993l-8.68-5.107l.091-10.054l8.772-4.947l8.68 5.106l-.092 10.054Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.521 19.052l8.369 4.917l8.778-4.787m-8.778 4.892l-.001 9.902m.073.312v8.035m8.901-23.19l9.598-5.621m-27.336 5.534l-9.54-5.527"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}