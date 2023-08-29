package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Genderless(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M192 144a112 112 0 1 1 0 224a112 112 0 1 1 0-224zm0 288a176 176 0 1 0 0-352a176 176 0 1 0 0 352z"/>`),
		g.Group(children),
	)
}