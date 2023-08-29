package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sitemap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.636 17.455V10.91H12V6.546h2.182V.001H7.637v6.545h2.182v4.364H2.183v6.545H.001V24h6.545v-6.545H4.364v-4.364h13.091v4.364h-2.182V24h6.545v-6.545z"/>`),
		g.Group(children),
	)
}