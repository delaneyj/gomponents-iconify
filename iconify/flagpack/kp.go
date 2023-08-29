package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackKp0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackKp2)"><use href="#flagpackKp0"/><path fill="#3D58DB" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackKp1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackKp1)"><path fill="#C51918" stroke="#F7FCFF" stroke-width="2" d="M0 5h-1v14h34V5H0Z"/><path fill="#F7FCFF" fill-rule="evenodd" d="M10 17a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z" clip-rule="evenodd"/><path fill="#C51918" fill-rule="evenodd" d="M9.995 13.87L7.28 15.76l.958-3.168L5.6 10.593l3.309-.067L9.995 7.4l1.087 3.126l3.308.067l-2.637 2l.958 3.167l-2.716-1.89Z" clip-rule="evenodd"/></g></g><defs><clipPath id="flagpackKp2"><use href="#flagpackKp0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}