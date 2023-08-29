package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmWatsonStudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 21c-2.8 0-5-2.2-5-5s2.2-5 5-5s5 2.2 5 5s-2.2 5-5 5zm0-8c-1.7 0-3 1.3-3 3s1.3 3 3 3s3-1.3 3-3s-1.3-3-3-3zm-11-2V6c0-1.1-.9-2-2-2h-2v2h2v5c0 2.1 1.1 3.9 2.7 5c-1.6 1.1-2.7 2.9-2.7 5v5h-2v2h2c1.1 0 2-.9 2-2v-5c0-2.2 1.8-4 4-4v-2c-2.2 0-4-1.8-4-4zM2 30v-6h6v6H2zm2-4v2h2v-2H4zm-2-7v-6h6v6H2zm2-4v2h2v-2H4zM2 8V2h6v6H2zm2-4v2h2V4H4z"/>`),
		g.Group(children),
	)
}