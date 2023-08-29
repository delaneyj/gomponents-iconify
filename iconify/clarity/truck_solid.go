package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TruckSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M30 12h-4V7a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h1V8h20v13.49A4.45 4.45 0 0 0 21.25 24h-6.82a4.5 4.5 0 0 0-4.17-2.76A4.38 4.38 0 1 0 14.72 26H21a4.48 4.48 0 0 0 8.91 0H34V16a4 4 0 0 0-4-4ZM10.26 28a2.38 2.38 0 1 1 0-4.75a2.38 2.38 0 1 1 0 4.75Zm15.17 0a2.38 2.38 0 1 1 2.5-2.37A2.44 2.44 0 0 1 25.42 28ZM32 17h-6v-3h4a2 2 0 0 1 2 2Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}