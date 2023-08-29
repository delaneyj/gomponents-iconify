package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldWarningThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 44H48a12 12 0 0 0-12 12v58.77c0 86.88 73.54 115.7 88.28 120.59a11.47 11.47 0 0 0 7.44 0c14.74-4.87 88.28-33.69 88.28-120.57V56a12 12 0 0 0-12-12Zm4 70.79c0 81.38-69 108.41-82.81 113a3.51 3.51 0 0 1-2.39 0C113 223.2 44 196.17 44 114.79V56a4 4 0 0 1 4-4h160a4 4 0 0 1 4 4ZM124 136V96a4 4 0 0 1 8 0v40a4 4 0 0 1-8 0Zm12 36a8 8 0 1 1-8-8a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}