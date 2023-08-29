package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerampUnlocker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.584 43.718A21.428 21.428 0 0 1 24 45.5C12.126 45.5 2.5 35.874 2.5 24S12.126 2.5 24 2.5S45.5 12.126 45.5 24a21.41 21.41 0 0 1-2.2 9.486"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.584 37.505A15.925 15.925 0 0 1 24 40c-8.837 0-16-7.163-16-16S15.163 8 24 8s16 7.163 16 16c0 1.93-.342 3.78-.968 5.493"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.024 24.024L18.555 33.5l.028-19l16.44 9.524Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M33.364 35.548H44.72c.43 0 .78.349.78.78v8.392c0 .43-.35.78-.78.78H33.364a.78.78 0 0 1-.78-.78v-8.393c0-.43.35-.78.78-.78h0Z"/><path d="M34.768 35.548v-1.79a4.27 4.27 0 0 1 8.541 0v1.79"/><circle cx="39.038" cy="40.524" r="1.656"/></g>`),
		g.Group(children),
	)
}