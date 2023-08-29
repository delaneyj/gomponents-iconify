package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Awake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 2h2v4.96h-2zm6.687 6.89l3.507-3.506l1.414 1.414l-3.507 3.507zM25.04 15H30v2h-4.96zm-3.347 8.104l1.414-1.414l3.507 3.507L25.2 26.61zM15 25.04h2V30h-2zm-9.604.162l3.508-3.507l1.414 1.414l-3.507 3.507zM2 15h4.96v2H2zm3.39-8.197l1.415-1.414l3.507 3.507l-1.414 1.414zM16 12a4 4 0 1 1-4 4a4 4 0 0 1 4-4m0-2a6 6 0 1 0 6 6a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}