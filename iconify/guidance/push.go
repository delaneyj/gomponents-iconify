package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Push(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M24 12h-8m3.429-4c0 .423.419 1.056.842 1.587a7.506 7.506 0 0 0 1.944 1.738c.56.342 1.239.67 1.785.67M19.428 16c0-.423.42-1.056.843-1.587a7.509 7.509 0 0 1 1.944-1.738c.56-.342 1.239-.67 1.785-.67M0 22.5h6a6.5 6.5 0 0 0 6.5-6.5V9.5H12A2.5 2.5 0 0 0 9.5 12v4V6A2.5 2.5 0 0 0 7 3.5h-.59M.5 16V3.5H1A2.5 2.5 0 0 1 3.5 6v6V1.5h.46A2.5 2.5 0 0 1 6.46 4v8"/>`),
		g.Group(children),
	)
}