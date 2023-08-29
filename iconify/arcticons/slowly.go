package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slowly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 39.5v-31m-4.36 16.88c-1.4 0-1.4-3.324-2.8-3.324s-1.4 3.324-2.801 3.324c-1.403 0-1.403-3.324-2.805-3.324c-1.401 0-1.401 3.324-2.802 3.324c-1.404 0-1.404-3.324-2.807-3.324c-1.407 0-1.407 3.324-2.813 3.324a.82.82 0 0 1-.272-.045m17.1 7.105c-1.4 0-1.4-3.324-2.8-3.324s-1.4 3.324-2.801 3.324c-1.403 0-1.403-3.324-2.805-3.324c-1.401 0-1.401 3.324-2.802 3.324c-1.404 0-1.404-3.324-2.807-3.324c-.93 0-1.246 1.455-1.76 2.44M19.64 39.5c-1.4 0-1.4-3.324-2.8-3.324s-1.4 3.324-2.8 3.324c-1.403 0-1.403-3.324-2.805-3.324c-1.212 0-1.376 2.487-2.304 3.159m21.067-25.837h9v9h-9z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.946 20.076a2.122 2.122 0 0 0 1.552.422h.423a1.249 1.249 0 0 0 1.247-1.25h0a1.249 1.249 0 0 0-1.247-1.25h-.847a1.249 1.249 0 0 1-1.247-1.25h0a1.249 1.249 0 0 1 1.247-1.25h.424a2.122 2.122 0 0 1 1.551.422"/>`),
		g.Group(children),
	)
}