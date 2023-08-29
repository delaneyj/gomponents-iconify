package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tarp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M576 128c0-35.3-28.7-64-64-64H64C28.7 64 0 92.7 0 128v256c0 35.3 28.7 64 64 64h352V320c0-17.7 14.3-32 32-32h128V128zM448 448l128-128H448v128zM96 128a32 32 0 1 1 0 64a32 32 0 1 1 0-64z"/>`),
		g.Group(children),
	)
}