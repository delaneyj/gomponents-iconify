package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandFistLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 82h-18V64a30 30 0 0 0-54-18a30 30 0 0 0-54 18a30 30 0 0 0-48 24v40a102 102 0 0 0 204 0v-16a30 30 0 0 0-30-30Zm-48-36a18 18 0 0 1 18 18v18h-34a14.71 14.71 0 0 0-2 .16V64a18 18 0 0 1 18-18ZM86 64a18 18 0 0 1 36 0v40a18 18 0 0 1-36 0ZM38 88a18 18 0 0 1 36 0v16a18 18 0 0 1-36 0Zm180 40a90 90 0 0 1-180 0a30 30 0 0 0 42-6a30 30 0 0 0 45.12 3.3A30.18 30.18 0 0 0 140 139.51A45.92 45.92 0 0 0 122 176a6 6 0 0 0 12 0a34 34 0 0 1 34-34a6 6 0 0 0 0-12h-16a18 18 0 0 1-18-18V96a2 2 0 0 1 2-2h64a18 18 0 0 1 18 18Z"/>`),
		g.Group(children),
	)
}