package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudMoonFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M164 72a76.35 76.35 0 0 0-12.36 1A71.93 71.93 0 0 0 96.17 9.83a8 8 0 0 0-9.59 9.58A56.45 56.45 0 0 1 88 32a56.06 56.06 0 0 1-56 56a56.45 56.45 0 0 1-12.59-1.42a8 8 0 0 0-9.59 9.59a72.22 72.22 0 0 0 32.29 45.06A52 52 0 0 0 84 224h80a76 76 0 0 0 0-152ZM29.37 104H32a72.08 72.08 0 0 0 72-72v-2.67a55.64 55.64 0 0 1 32 48.05a76.4 76.4 0 0 0-43 43.4a52.38 52.38 0 0 0-9-.78a51.69 51.69 0 0 0-30 9.59A56.22 56.22 0 0 1 29.37 104Z"/>`),
		g.Group(children),
	)
}