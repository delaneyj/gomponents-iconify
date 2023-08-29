package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airtel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.67 24.09C11.5 32.74 24.7 17 32.32 17c3 0 3.16 2.41 3.16 4.92c0 7.16-15.16 20.06-17.66 17.29S26.6 32.83 21.57 29c-3.63-2.75-9.55 5.15-9.14 9.13c.3 2.89 1.1 4.37 5 4.37c8.11 0 24.83-9 24.54-25.47c-.1-6.72-4.26-11.53-12.62-11.53S2.27 16.22 6.67 24.09Z"/>`),
		g.Group(children),
	)
}