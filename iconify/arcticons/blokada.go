package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blokada(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.13 41.3A17.63 17.63 0 0 0 24 43.5c.44 0 17.2-12.5 17.2-28.27c0-5.41-1.2-10.11-1.42-10.11c-.62 0-1.78 2.57-6 2.57c-5.35 0-8.72-3.19-9.78-3.19s-4.43 3.19-9.75 3.19c-4.25 0-5.41-2.57-6-2.57c-.25 0-1.45 4.7-1.45 10.11c0 11.83 9.43 21.82 14.33 26.07m12.62-30.24a10.65 10.65 0 0 0 3.72-.62a42.59 42.59 0 0 1 .27 4.79a22.43 22.43 0 0 1-.89 5.85L21.78 9.29a11.31 11.31 0 0 0 2.13-1.16a20 20 0 0 0 9.84 2.93Zm-23.58 4.16a31.38 31.38 0 0 1 .27-4.16L34.37 27l-24.11-9.83a14.07 14.07 0 0 1-.09-1.95ZM24 39.33a50.79 50.79 0 0 1-6.83-7C14.6 29.05 11.5 24.27 10.53 19L28.7 34.81a37.63 37.63 0 0 1-4.7 4.52Z"/>`),
		g.Group(children),
	)
}