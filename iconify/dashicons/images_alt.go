package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagesAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 15v-3H2V2h12v3h2v3h2v10H6v-3H4zm7-12c-1.1 0-2 .9-2 2h4a2 2 0 0 0-2-2zm-7 8V6H3v5h1zm7-3h4a2 2 0 1 0-4 0zm-5 6V9H5v5h1zm9-1a2 2 0 1 0 .001-3.999A2 2 0 0 0 15 13zm2 4v-2c-5 0-5-3-10-3v5h10z"/>`),
		g.Group(children),
	)
}