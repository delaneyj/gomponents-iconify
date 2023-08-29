package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rexx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 512V0h188.137l141.037 200.838L201 287l152 225H242.445L127 346V235l78.887-51.734L127 78H88v434z"/>`),
		g.Group(children),
	)
}