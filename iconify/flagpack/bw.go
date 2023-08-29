package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackBw0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackBw2)"><use href="#flagpackBw0"/><path fill="#42ADDF" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackBw1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackBw1)"><path fill="#58A5FF" fill-rule="evenodd" d="M0 0v8h32V0H0Z" clip-rule="evenodd"/><path fill="#272727" stroke="#F7FCFF" stroke-width="2" d="M0 9h-1v6h34V9H0Z"/></g></g><defs><clipPath id="flagpackBw2"><use href="#flagpackBw0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}