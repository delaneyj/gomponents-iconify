package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 12C12.123 12 2 32 2 32s10.123 20 30 20c19.879 0 30-20 30-20S51.879 12 32 12zm0 36.664C15.436 48.664 7 32 7 32s8.436-16.668 25-16.668C48.566 15.332 57 32 57 32s-8.434 16.664-25 16.664z"/><path fill="currentColor" d="M31.9 46.5c-7.995 0-14.5-6.505-14.5-14.5s6.505-14.5 14.5-14.5S46.4 24.005 46.4 32s-6.504 14.5-14.5 14.5m0-26c-6.341 0-11.5 5.159-11.5 11.5s5.159 11.5 11.5 11.5S43.4 38.341 43.4 32s-5.159-11.5-11.5-11.5"/><path fill="currentColor" d="M39.398 31.994c0 4.148-3.359 7.51-7.496 7.51a7.505 7.505 0 0 1-7.504-7.51c0-4.141 3.358-7.49 7.504-7.49c4.137 0 7.496 3.35 7.496 7.49"/>`),
		g.Group(children),
	)
}