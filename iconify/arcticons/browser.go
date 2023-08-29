package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Browser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.51 21.51 0 0 0 2.5 24h0A21.51 21.51 0 0 0 24 45.5h0A21.51 21.51 0 0 0 45.5 24h0A21.51 21.51 0 0 0 24 2.5Zm-4.46 4.1L22 9.29l-1.52 2.85l-5 2.92l-1.88 6.42l8.61 2.73L13 38.19a18.08 18.08 0 0 1-2.75-2.64L8.87 24l4.73-2.49l-6.73-2.95a18 18 0 0 1 12.66-12ZM32 7.92a17.88 17.88 0 0 1 5.74 4.51l-8.74.32Zm6.76 5.84a17.91 17.91 0 0 1-3 23.82l-.92-10.21l-6-1.7l-1.39-6.8Z"/>`),
		g.Group(children),
	)
}