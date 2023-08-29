package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SirenLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M122 16V8a6 6 0 0 1 12 0v8a6 6 0 0 1-12 0Zm78 30a6 6 0 0 0 4.24-1.76l8-8a6 6 0 1 0-8.48-8.48l-8 8A6 6 0 0 0 200 46ZM51.76 44.24a6 6 0 0 0 8.48-8.48l-8-8a6 6 0 0 0-8.48 8.48ZM137 74.08a6 6 0 1 0-2 11.84c20 3.34 35 21.44 35 42.08a6 6 0 0 0 12 0c0-26.43-19.35-49.61-45-53.92ZM230 176v24a14 14 0 0 1-14 14H40a14 14 0 0 1-14-14v-24a14 14 0 0 1 14-14h2v-34a86 86 0 0 1 86-86h.65c47.06.35 85.35 39.38 85.35 87v33h2a14 14 0 0 1 14 14ZM54 162h148v-33c0-41-32.94-74.7-73.44-75H128a74 74 0 0 0-74 74Zm164 14a2 2 0 0 0-2-2H40a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h176a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}