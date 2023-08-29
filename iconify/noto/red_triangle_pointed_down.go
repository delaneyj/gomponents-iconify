package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedTrianglePointedDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#C33" d="m61.32 90l-28-48.5a3.68 3.68 0 0 1 3.3-5.4h56c2.03-.03 3.7 1.6 3.73 3.63c.01.62-.14 1.22-.43 1.77l-28 48.5a3.642 3.642 0 0 1-4.84 1.76c-.78-.35-1.4-.98-1.76-1.76z"/><path fill="#F44336" d="M58.59 85.3L33.09 41a3.378 3.378 0 0 1 3-4.9h51a3.348 3.348 0 0 1 3.37 3.35c0 .54-.13 1.07-.37 1.55l-25.4 44.3a3.418 3.418 0 0 1-6.1 0z"/><path fill="#FF8A80" d="M44.6 43.58c1.02 1.14 18.79 30.98 18.79 30.98s1.52 2.29-.25 3.68c-1.78 1.4-3.43 0-4.19-1.14S42.06 49.17 41.43 47.9a4.45 4.45 0 0 1 .63-4.32c.89-.89 1.91-.89 2.54 0z"/>`),
		g.Group(children),
	)
}