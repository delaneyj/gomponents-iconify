package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandGooglePodcasts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v2m0 14v2m0-13v8m-4 1v2m-4-8v2m16-2v2M8 5v8m8-6V5m0 14v-8"/>`),
		g.Group(children),
	)
}