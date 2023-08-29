package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M107 3q88 0 150.5 62.5T320 216t-62.5 150.5T107 429q-57 0-107-28q49-29 78-78t29-107t-29-107T0 31Q50 3 107 3z"/>`),
		g.Group(children),
	)
}