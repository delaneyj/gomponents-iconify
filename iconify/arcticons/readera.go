package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Readera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.6 13.09h7.29m-7.34 3.12h7.3m-7.3 3.07h7.3m-7.3 3.15h7.3M8.67 9v25.08h11.12c1.79 0 4.29 1.75 4.29 1.75L24 10.35a5.51 5.51 0 0 0-3.85-1.58Zm22.15 25.08h-2.45c-1.79 0-4.29 1.75-4.29 1.75l.09-25.48A5.51 5.51 0 0 1 28 8.77h2.8m3.09.09l5.6.09v25.13h-5.6"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.89 36.62h7.7V10.28H40m-31.51.09h-2.2V36.6h15.16a10.39 10.39 0 0 0 2.53.58a9.37 9.37 0 0 0 2.38-.56h4.46m0-29.43v32.4l1.49-1.05l1.58 1V7.28Z"/>`),
		g.Group(children),
	)
}