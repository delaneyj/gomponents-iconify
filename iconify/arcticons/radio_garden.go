package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioGarden(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5C12.128 2.505 2.506 12.128 2.5 24h0c.005 11.872 9.628 21.494 21.5 21.5h0c11.872-.005 21.494-9.628 21.5-21.5h0C45.495 12.128 35.872 2.506 24 2.5Zm-4.46 4.1L22 9.29l-1.52 2.85l-5 2.92l-1.88 6.42l8.61 2.73L13 38.19a18.084 18.084 0 0 1-2.75-2.64L8.87 24l4.73-2.49l-6.73-2.95a18 18 0 0 1 12.66-12l.01.04ZM32 7.92a17.88 17.88 0 0 1 5.74 4.51l-8.74.32l3-4.83Zm6.76 5.84c5.258 7.534 3.961 17.825-3 23.82l-.92-10.21l-6-1.7l-1.39-6.8l11.31-5.11Z"/><circle cx="24" cy="24.023" r="8.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}