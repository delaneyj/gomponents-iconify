package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 48a208 208 0 1 1 0 416a208 208 0 1 1 0-416zm0 464a256 256 0 1 0 0-512a256 256 0 1 0 0 512zM135.1 217.4c-4.5 4.2-7.1 10.1-7.1 16.3c0 12.3 10 22.3 22.3 22.3H208v96c0 17.7 14.3 32 32 32h32c17.7 0 32-14.3 32-32v-96h57.7c12.3 0 22.3-10 22.3-22.3c0-6.2-2.6-12.1-7.1-16.3l-107.1-99.9c-3.8-3.5-8.7-5.5-13.8-5.5s-10.1 2-13.8 5.5l-107.1 99.9z"/>`),
		g.Group(children),
	)
}