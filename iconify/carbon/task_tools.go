package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskTools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28.5 32a3.476 3.476 0 0 1-2.475-1.025l-4.128-4.128a6.496 6.496 0 0 1-7.348-8.956l.588-1.34l4.156 4.156a1.023 1.023 0 0 0 1.413 0a.999.999 0 0 0 .001-1.414l-4.156-4.157l1.34-.588a6.497 6.497 0 0 1 8.956 7.349l4.128 4.128A3.5 3.5 0 0 1 28.5 32Zm-6.03-7.409l4.97 4.97a1.535 1.535 0 0 0 2.12 0a1.498 1.498 0 0 0 0-2.121l-4.969-4.97l.188-.583A4.496 4.496 0 0 0 20.5 16q-.126 0-.25.007l1.872 1.872a3 3 0 0 1 0 4.242a3.072 3.072 0 0 1-4.243 0l-1.872-1.872Q16 20.374 16 20.5a4.497 4.497 0 0 0 5.888 4.28Z"/><path fill="currentColor" d="M25 5h-3V4a2.006 2.006 0 0 0-2-2h-8a2.006 2.006 0 0 0-2 2v1H7a2.006 2.006 0 0 0-2 2v21a2.006 2.006 0 0 0 2 2h7v-2H7V7h3v3h12V7h3v5h2V7a2.006 2.006 0 0 0-2-2Zm-5 3h-8V4h8Z"/>`),
		g.Group(children),
	)
}