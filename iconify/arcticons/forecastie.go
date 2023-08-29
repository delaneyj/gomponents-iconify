package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forecastie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.41 11c5.3-.29 10.8 2.76 12.63 9.75c10.38 1.41 9.49 15.3.06 16.27H13.28c-10.95-1.47-12.44-17.4 0-19.54A12 12 0 0 1 23.41 11Zm0 0"/>`),
		g.Group(children),
	)
}