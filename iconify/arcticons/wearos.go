package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wearos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.5 40.43a4.34 4.34 0 0 1-1.71.36A4.19 4.19 0 0 1 14 38.3L4.21 16.4a4.19 4.19 0 1 1 7.66-3.4l9.75 21.9a4.19 4.19 0 0 1-2.12 5.53Zm11.65 0a4.34 4.34 0 0 1-1.71.36a4.19 4.19 0 0 1-3.83-2.49l-9.75-21.9a4.19 4.19 0 1 1 7.66-3.4l9.75 21.9a4.19 4.19 0 0 1-2.12 5.53Zm4.3-13.2A4.19 4.19 0 1 1 39.64 23a4.18 4.18 0 0 1-4.19 4.23Zm3.7-8.84a3.7 3.7 0 1 1 3.7-3.7a3.7 3.7 0 0 1-3.7 3.7Z"/>`),
		g.Group(children),
	)
}