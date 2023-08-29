package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M116 128V48a12 12 0 0 1 24 0v80a12 12 0 0 1-24 0Zm66.55-82a12 12 0 0 0-13.1 20.1C191.41 80.37 204 103 204 128a76 76 0 0 1-152 0c0-25 12.59-47.63 34.55-61.95A12 12 0 0 0 73.45 46C44.56 64.78 28 94.69 28 128a100 100 0 0 0 200 0c0-33.31-16.56-63.22-45.45-82Z"/>`),
		g.Group(children),
	)
}