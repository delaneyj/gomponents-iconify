package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eapps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.38 19.84h-7.86l-7.77-13A2 2 0 0 0 23 6.06a2 2 0 0 0-.78.78l-7.77 13H6.62A3.13 3.13 0 0 0 3.51 23a3.09 3.09 0 0 0 .11.8L8 40.08a2.84 2.84 0 0 0 2.74 2.1h26.5a2.84 2.84 0 0 0 2.76-2.1l4.41-16.31a3.11 3.11 0 0 0-2.2-3.77a2.91 2.91 0 0 0-.83-.16Zm-17.38-8l4.77 8h-9.54Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.51 37.01l-6.51-12l-6.51 12M19.67 33h8.62"/>`),
		g.Group(children),
	)
}