package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kanadrill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.9 5.83C23.15 16 21.19 24.64 24.32 42.17"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.81 12.49C18.27 14 24.85 14.8 41 9m-7.77 7C31.48 22.6 23.59 40 11 40C-.83 40 7.3 23.12 27.05 21.45c21.94-1.86 17 12.32 6.78 20"/>`),
		g.Group(children),
	)
}