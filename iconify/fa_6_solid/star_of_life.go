package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOfLife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M208 32c0-17.7 14.3-32 32-32h32c17.7 0 32 14.3 32 32v140.9l122-70.4c15.3-8.8 34.9-3.6 43.7 11.7l16 27.7c8.8 15.3 3.6 34.9-11.7 43.7L352 256l122 70.4c15.3 8.8 20.5 28.4 11.7 43.7l-16 27.7c-8.8 15.3-28.4 20.6-43.7 11.7l-122-70.4V480c0 17.7-14.3 32-32 32h-32c-17.7 0-32-14.3-32-32V339.1L86 409.6c-15.3 8.8-34.9 3.6-43.7-11.7l-16-27.7c-8.8-15.3-3.6-34.9 11.7-43.7L160 256L38 185.6c-15.3-8.8-20.5-28.4-11.7-43.7l16-27.7c8.8-15.4 28.4-20.6 43.7-11.8l122 70.4V32z"/>`),
		g.Group(children),
	)
}