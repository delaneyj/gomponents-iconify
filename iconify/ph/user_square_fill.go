package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M172 120a44 44 0 1 1-44-44a44 44 0 0 1 44 44Zm52-72v160a16 16 0 0 1-16 16H48a16 16 0 0 1-16-16V48a16 16 0 0 1 16-16h160a16 16 0 0 1 16 16Zm-16 160V48H48v160h3.67a80.58 80.58 0 0 1 26.07-38.25q3.08-2.48 6.36-4.62a4 4 0 0 1 4.81.33a59.82 59.82 0 0 0 78.18 0a4 4 0 0 1 4.81-.33q3.28 2.15 6.36 4.62A80.58 80.58 0 0 1 204.33 208H208Z"/>`),
		g.Group(children),
	)
}