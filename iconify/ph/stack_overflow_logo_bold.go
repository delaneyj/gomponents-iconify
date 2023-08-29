package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackOverflowLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 151.67V216a12 12 0 0 1-12 12H48a12 12 0 0 1-12-12v-64.33a12 12 0 1 1 24 0v52.23h136v-52.23a12 12 0 1 1 24 0ZM88 183.81h80a12.06 12.06 0 0 0 0-24.11H88a12.06 12.06 0 0 0 0 24.11ZM96.2 113l75.17 27.49a12.05 12.05 0 0 0 8.21-22.66l-75.17-27.48A12 12 0 0 0 96.2 113ZM128 49.29l61.29 51.64a12 12 0 0 0 16.9-1.48a12.09 12.09 0 0 0-1.48-17l-61.27-51.63a12 12 0 0 0-16.91 1.49A12.1 12.1 0 0 0 128 49.29Z"/>`),
		g.Group(children),
	)
}