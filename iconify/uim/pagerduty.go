package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pagerduty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 16.673h2.93V22H6zM17.034 2.977C15.46 2.139 14.37 2 11.796 2H6v12.124h5.77c2.296 0 4.008-.14 5.517-1.141a5.769 5.769 0 0 0 2.499-4.997a5.487 5.487 0 0 0-2.752-5.01Zm-4.591 8.61H8.93V4.6l3.31-.026c3.018-.038 4.527 1.028 4.527 3.437c0 2.588-1.864 3.577-4.324 3.577Z"/>`),
		g.Group(children),
	)
}