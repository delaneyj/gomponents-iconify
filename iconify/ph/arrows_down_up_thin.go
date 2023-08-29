package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsDownUpThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M114.83 173.17a4 4 0 0 1 0 5.66l-32 32a4 4 0 0 1-5.66 0l-32-32a4 4 0 0 1 5.66-5.66L76 198.34V48a4 4 0 0 1 8 0v150.34l25.17-25.17a4 4 0 0 1 5.66 0Zm96-96l-32-32a4 4 0 0 0-5.66 0l-32 32a4 4 0 0 0 5.66 5.66L172 57.66V208a4 4 0 0 0 8 0V57.66l25.17 25.17a4 4 0 1 0 5.66-5.66Z"/>`),
		g.Group(children),
	)
}