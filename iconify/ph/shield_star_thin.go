package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldStarThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M171.71 118.52a4 4 0 0 1-2.22 5.19l-35.16 14.07l20.87 27.82a4 4 0 1 1-6.4 4.8L128 142.67l-20.8 27.73a4 4 0 0 1-6.4-4.8l20.87-27.82l-35.16-14.07a4 4 0 0 1 3-7.42l34.49 13.8V96a4 4 0 0 1 8 0v34.09l34.51-13.8a4 4 0 0 1 5.2 2.23ZM220 56v58.77c0 86.88-73.54 115.7-88.28 120.59a11.47 11.47 0 0 1-7.44 0C109.54 230.49 36 201.67 36 114.79V56a12 12 0 0 1 12-12h160a12 12 0 0 1 12 12Zm-8 0a4 4 0 0 0-4-4H48a4 4 0 0 0-4 4v58.77c0 81.38 69 108.41 82.8 113a3.51 3.51 0 0 0 2.39 0C143 223.2 212 196.17 212 114.79Z"/>`),
		g.Group(children),
	)
}