package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSlashThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M51 37.31a4 4 0 0 0-6 5.38l41.16 45.24q-1.38 2.55-2.59 5.19A60 60 0 1 0 72 212h88a83.19 83.19 0 0 0 32.88-6.69L205 218.69a4 4 0 1 0 5.92-5.38ZM160 204H72a52 52 0 0 1 0-104a52.92 52.92 0 0 1 8.54.72A84.21 84.21 0 0 0 76 128a4 4 0 0 0 8 0a76 76 0 0 1 7.9-33.76L187.13 199a75.37 75.37 0 0 1-27.13 5Zm84-76a83.86 83.86 0 0 1-21.34 55.94a4 4 0 1 1-6-5.33A76 76 0 0 0 115 66.75a4 4 0 0 1-4.74-6.45A84 84 0 0 1 244 128Z"/>`),
		g.Group(children),
	)
}