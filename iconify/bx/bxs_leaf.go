package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsLeaf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M22 3.41l-.12-1.26l-1.2.4a13.84 13.84 0 0 1-6.41.64a11.87 11.87 0 0 0-6.68.9A7.23 7.23 0 0 0 3.3 9.5a9 9 0 0 0 .39 4.58a16.6 16.6 0 0 1 1.18-2.2a9.85 9.85 0 0 1 4.07-3.43a11.16 11.16 0 0 1 5.06-1A12.08 12.08 0 0 0 9.34 9.2a9.48 9.48 0 0 0-1.86 1.53a11.38 11.38 0 0 0-1.39 1.91a16.39 16.39 0 0 0-1.57 4.54A26.42 26.42 0 0 0 4 22h2a30.69 30.69 0 0 1 .59-4.32a9.25 9.25 0 0 0 4.52 1.11a11 11 0 0 0 4.28-.87C23 14.67 22 3.86 22 3.41z" fill="currentColor"/>`),
		g.Group(children),
	)
}