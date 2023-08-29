package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Is(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackIs0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackIs2)"><use href="#flagpackIs0"/><path fill="#0061C1" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackIs1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackIs1)"><path fill="#D21034" stroke="#F1F9FF" stroke-width="2" d="M10-1H9V9H-1v6H9v10h6V15h18V9H15V-1h-5Z"/></g></g><defs><clipPath id="flagpackIs2"><use href="#flagpackIs0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}