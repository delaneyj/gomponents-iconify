package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M24.73 28.146a.498.498 0 0 0 .36-.153l7.359-7.631a.498.498 0 0 0 0-.693l-7.391-7.662a.5.5 0 1 0-.72.693l7.056 7.315l-7.024 7.284a.5.5 0 0 0 .36.847z"/><path fill="currentColor" d="M7.91 20.515h24.18a.5.5 0 0 0 0-1H7.91a.5.5 0 0 0 0 1z"/>`),
		g.Group(children),
	)
}