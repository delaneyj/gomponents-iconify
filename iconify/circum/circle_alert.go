package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 9a.5.5 0 0 0-1 0v4.02a.5.5 0 0 0 1 0Z"/><circle cx="12" cy="15.001" r=".5" fill="currentColor"/><path fill="currentColor" d="M12 21.935A9.933 9.933 0 1 1 21.934 12A9.945 9.945 0 0 1 12 21.935Zm0-18.866A8.933 8.933 0 1 0 20.934 12A8.944 8.944 0 0 0 12 3.069Z"/>`),
		g.Group(children),
	)
}