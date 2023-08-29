package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsNineFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16ZM76 192a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-52a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-52a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm52 104a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-52a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-52a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm52 104a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-52a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-52a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}