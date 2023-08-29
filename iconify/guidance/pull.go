package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M11 5h3v.5a3 3 0 0 1-3 3H7h5.25a1.75 1.75 0 1 1 0 3.5H10h1.25a1.75 1.75 0 1 1 0 3.5H10h.25a1.75 1.75 0 1 1 0 3.5H0M8 5H3.5A3.5 3.5 0 0 1 0 8.5m9.5 0v-8h5M9.5 19v4.5h5M16 12h8m-3.429-4c0 .423-.419 1.056-.842 1.587a7.506 7.506 0 0 1-1.944 1.738c-.56.342-1.239.67-1.785.67M20.572 16c0-.423-.42-1.056-.843-1.587a7.509 7.509 0 0 0-1.944-1.738c-.56-.342-1.239-.67-1.785-.67"/>`),
		g.Group(children),
	)
}