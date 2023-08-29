package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PearlOfTheOrient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.997.999L13 3.17a3.001 3.001 0 0 1 0 5.66v4.296a4.002 4.002 0 0 1 1.86 6.67L16.753 23h-2.325l-1.28-2.168a3.998 3.998 0 0 1-2.298 0L9.573 23H7.249l1.89-3.204A3.906 3.906 0 0 1 8 17a4.002 4.002 0 0 1 3-3.874V8.829a3.001 3.001 0 0 1 0-5.658L10.997 1l2-.002Zm-.997 4a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm0 10a2 2 0 0 0-2 2c0 .518.165.945.502 1.327A1.985 1.985 0 0 0 12 19a1.986 1.986 0 0 0 1.525-.706A2 2 0 0 0 12 15Z"/>`),
		g.Group(children),
	)
}