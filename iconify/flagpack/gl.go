package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackGl0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackGl3)"><use href="#flagpackGl0"/><path fill="#C51918" fill-rule="evenodd" d="M0 12h32v12H0V12Z" clip-rule="evenodd"/><mask id="flagpackGl1" width="32" height="12" x="0" y="12" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 12h32v12H0V12Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackGl1)"><path fill="#F7FCFF" fill-rule="evenodd" d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z" clip-rule="evenodd"/></g><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0h32v12H0V0Z" clip-rule="evenodd"/><mask id="flagpackGl2" width="32" height="12" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0h32v12H0V0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackGl2)"><path fill="#C51918" fill-rule="evenodd" d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z" clip-rule="evenodd"/></g></g><defs><clipPath id="flagpackGl3"><use href="#flagpackGl0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}