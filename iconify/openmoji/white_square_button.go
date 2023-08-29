package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteSquareButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M46 26H25v21h21V26Z"/><path fill="#fff" d="M56 16H16v40h40V16Z"/><path fill="#3F3F3F" d="M46 26H25v21h21V26Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M46 26H25v21h21V26Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M56 16H16v40h40V16Z"/>`),
		g.Group(children),
	)
}