package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackSquareButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill-rule="evenodd" d="M16 16h40v40H16V16Zm9 10v21h21V26H25Z" clip-rule="evenodd"/><path fill="#3F3F3F" fill-rule="evenodd" d="M16 16h40v40H16V16Zm9 10h21v21H25V26Z" clip-rule="evenodd"/><path fill="#fff" d="M46 26H25v21h21V26Z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M56 16H16v40h40V16Z"/><path d="M46 26H25v21h21V26Z"/></g>`),
		g.Group(children),
	)
}