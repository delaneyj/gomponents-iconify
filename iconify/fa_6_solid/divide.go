package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Divide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M272 96a48 48 0 1 0-96 0a48 48 0 1 0 96 0zm0 320a48 48 0 1 0-96 0a48 48 0 1 0 96 0zm128-128c17.7 0 32-14.3 32-32s-14.3-32-32-32H48c-17.7 0-32 14.3-32 32s14.3 32 32 32h352z"/>`),
		g.Group(children),
	)
}