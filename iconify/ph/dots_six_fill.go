package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsSixFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 48H32a16 16 0 0 0-16 16v128a16 16 0 0 0 16 16h192a16 16 0 0 0 16-16V64a16 16 0 0 0-16-16ZM68 168a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-56a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm60 56a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-56a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm60 56a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-56a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}