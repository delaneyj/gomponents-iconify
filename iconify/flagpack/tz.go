package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackTz0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackTz2)"><use href="#flagpackTz0"/><path fill="#3195F9" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackTz1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackTz1)"><path fill="#73BE4A" fill-rule="evenodd" d="M0 0v24L32 0H0Z" clip-rule="evenodd"/><path fill="#272727" stroke="#FFD018" stroke-width="2.5" d="m-1.822 25.44l.694 1.039l1.04-.694L36.172 1.58l1.04-.693l-.693-1.04l-2.221-3.327l-.694-1.04l-1.04.694l-36.26 24.204l-1.04.694l.694 1.04l2.22 3.326Z"/></g></g><defs><clipPath id="flagpackTz2"><use href="#flagpackTz0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}