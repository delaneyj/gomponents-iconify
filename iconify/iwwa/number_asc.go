package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberAsc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M21.219 8.008h2.644a.75.75 0 0 0 0-1.5h-2.644a.75.75 0 0 0 0 1.5zm0 6.414h5.551a.75.75 0 0 0 0-1.5h-5.551a.75.75 0 0 0 0 1.5zm0 12.828h13.459a.75.75 0 0 0 0-1.5H21.219a.75.75 0 0 0 0 1.5zM38 32.164H21.219a.75.75 0 0 0 0 1.5H38a.75.75 0 0 0 0-1.5zM21.219 20.836h9.459a.75.75 0 0 0 0-1.5h-9.459a.75.75 0 0 0 0 1.5zm-3.751 3.619l-6.692 8.076V6.934a.75.75 0 0 0-1.5 0v25.563l-6.693-8.043a.75.75 0 1 0-1.154.959l8.037 9.657a.75.75 0 0 0 .577.271h.001c.223 0 .434-.1.577-.271l8.003-9.657a.75.75 0 0 0-1.156-.958z"/>`),
		g.Group(children),
	)
}