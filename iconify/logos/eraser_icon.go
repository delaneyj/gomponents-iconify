package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EraserIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#EB043B" d="M121.5 97.041H7.542c-5.844 0-9.466-8.093-6.461-14.424l37.08-77.97C39.56 1.764 41.989 0 44.664 0h76.838v97.041Z"/><path fill="#0085FF" d="M134.54 0h113.92c5.843 0 9.464 8.093 6.46 14.424l-37.081 77.97c-1.357 2.844-3.826 4.607-6.461 4.607H134.54V0Z"/>`),
		g.Group(children),
	)
}