package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Orbot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.88 4.5h0l.21.17q4 3.11 2.69 7.33q-4.37-1.8-2.95-7.33a.59.59 0 0 1 0-.17Zm6.62 3.84c-.24 2-1.23 3.26-2.95 3.66c-.24-1.85.74-3.08 2.95-3.66Zm-7.09 3.77h0a28.09 28.09 0 0 0 3.76.62h0a12.6 12.6 0 0 0 3.39-.17A4.8 4.8 0 0 0 30.34 17a13.87 13.87 0 0 1 3.53 2.6A13.48 13.48 0 0 1 38 29.53h0a13.42 13.42 0 0 1-4.09 9.87a13.41 13.41 0 0 1-9.69 4.09H24a13.43 13.43 0 0 1-9.87-4.09A13.44 13.44 0 0 1 10 29.52a13.45 13.45 0 0 1 4.09-9.87a14.18 14.18 0 0 1 3.39-2.54l.41-.21a4.7 4.7 0 0 0 2.49-4.18Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.7 40.1a10.17 10.17 0 0 1 0-20.34h0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.7 36.19a6.13 6.13 0 0 1 0-12.25h0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.7 32.78a2.62 2.62 0 0 1 0-5.22h0"/>`),
		g.Group(children),
	)
}