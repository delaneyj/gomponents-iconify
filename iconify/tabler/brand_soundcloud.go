package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandSoundcloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 11h1c1.38 0 3 1.274 3 3c0 1.657-1.5 3-3 3h-6V7c3 0 4.5 1.5 5 4zM9 8v9m-3 0v-7m-3 6v-2"/>`),
		g.Group(children),
	)
}