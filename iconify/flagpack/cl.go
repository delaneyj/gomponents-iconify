package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackCl0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackCl1)"><use href="#flagpackCl0"/><path fill="#3D58DB" fill-rule="evenodd" d="M0 0h14v14H0V0Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="M14-2h18v16H14V-2Z" clip-rule="evenodd"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 14h32v10H0V14Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="m7.014 9.783l-3.803 2.272L5.088 8.19L1.764 5.68L5.6 5.639L7.05 2.21l.932 3.428l3.632.017L8.851 8.11l1.434 3.944l-3.27-2.272Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackCl1"><use href="#flagpackCl0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}