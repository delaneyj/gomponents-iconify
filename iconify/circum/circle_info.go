package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleInfo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 15a.5.5 0 0 0 1 0v-4.019a.5.5 0 0 0-1 0Z"/><circle cx="12" cy="8.999" r=".5" fill="currentColor"/><path fill="currentColor" d="M12 2.065A9.934 9.934 0 1 1 2.066 12A9.945 9.945 0 0 1 12 2.065Zm0 18.867A8.934 8.934 0 1 0 3.066 12A8.944 8.944 0 0 0 12 20.932Z"/>`),
		g.Group(children),
	)
}