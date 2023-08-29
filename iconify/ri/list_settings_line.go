package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListSettingsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 18h7v2H2v-2Zm0-7h9v2H2v-2Zm0-7h20v2H2V4Zm18.674 9.025l1.156-.391l1 1.732l-.916.805a4.014 4.014 0 0 1 0 1.658l.916.805l-1 1.732l-1.156-.391a4 4 0 0 1-1.435.83L19 21h-2l-.24-1.196a3.997 3.997 0 0 1-1.434-.83l-1.156.392l-1-1.732l.916-.805a4.014 4.014 0 0 1 0-1.658l-.916-.805l1-1.732l1.156.391c.41-.37.898-.655 1.435-.83L17 11h2l.24 1.196a3.99 3.99 0 0 1 1.434.83ZM18 18a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`),
		g.Group(children),
	)
}