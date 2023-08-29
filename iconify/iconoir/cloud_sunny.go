package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSunny(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 13c-1.667 0-5 1-5 5s3.333 5 5 5h12c1.667 0 5-1 5-5s-3.333-5-5-5m-6-1a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm7-3h1m-8-7V1m6.5 2.5l-1 1m-12-1l1 1M4 9h1"/>`),
		g.Group(children),
	)
}