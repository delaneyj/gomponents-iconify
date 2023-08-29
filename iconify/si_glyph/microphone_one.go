package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="m16.608 13.715l-5.567-6.416c-.488 1.203-1.523 2.176-2.705 2.723l6.416 5.525a1.323 1.323 0 0 0 1.856 0a1.281 1.281 0 0 0 0-1.832zM8.736 1.49c-.42 1.39-1.244 2.744-2.389 3.89c-1.14 1.14-2.51 1.96-3.891 2.38a4.455 4.455 0 0 0 3.012 1.178c2.469 0 4.469-1.986 4.469-4.438a4.388 4.388 0 0 0-1.201-3.01z"/><path d="M7.877.771A4.46 4.46 0 0 0 5.469.062C3 .062 1 2.048 1 4.5c0 .897.273 1.729.733 2.428c1.364-.314 2.759-1.106 3.907-2.255C6.79 3.523 7.56 2.15 7.877.771z"/></g>`),
		g.Group(children),
	)
}