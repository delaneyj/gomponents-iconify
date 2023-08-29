package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cheese(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.783 3.024a1 1 0 0 1 .664.082l10 5A1 1 0 0 1 22 9v4a1 1 0 0 1-1 1c-.173 0-.456.06-.666.211c-.159.115-.334.315-.334.789c0 .474.175.674.334.789c.21.15.493.211.666.211a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H4.5a1 1 0 0 1-.894-.553c-1.34-2.679-2.02-6.427-1.136-9.824c.903-3.475 3.434-6.515 8.313-7.6zM8 17a1 1 0 1 1 0-2h.001a1 1 0 1 1 0 2H8zm4-4a1 1 0 0 0 1 1h.001a1 1 0 1 0 0-2H13a1 1 0 0 0-1 1zm3 4a1 1 0 1 1 0-2h.001a1 1 0 1 1 0 2H15z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}