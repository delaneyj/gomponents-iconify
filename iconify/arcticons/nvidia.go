package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nvidia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.16 25.46c-4.48 5.27-11.48 9.85-20.22 9.85c-15.87 0-18.78-14.14-18.78-14.14a22.83 22.83 0 0 1 17.73-8.47c10.33 0 17 8.47 17 8.47s-7.72 9.38-15.76 9.38c-8.66 0-10.92-8.61-10.92-8.61a14.48 14.48 0 0 1 10.34-4.37c4.75 0 8.82 4.95 8.82 4.95"/>`),
		g.Group(children),
	)
}