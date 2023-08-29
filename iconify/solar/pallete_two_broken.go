package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PalleteTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M7 3.341A9.934 9.934 0 0 1 12 2c5.523 0 10 4.489 10 10.026c0 8.152-8.162 2.393-9.738 4.9c-.395.628.032 1.41.555 1.935a1.68 1.68 0 0 1 0 2.372c-.523.525-1.235.838-1.97.753C5.867 21.413 2 17.172 2 12.026A10 10 0 0 1 3.345 7"/><circle cx="17.5" cy="11.5" r="1.5"/><circle cx="6.5" cy="11.5" r="1.5"/><path d="M11.085 7a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM16 7a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/></g>`),
		g.Group(children),
	)
}