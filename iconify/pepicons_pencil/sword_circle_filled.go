package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwordCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilSwordCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M10.912 15.084a.5.5 0 0 1-.002-.707l3.53-3.542a.5.5 0 0 1 .708.705l-3.53 3.543a.5.5 0 0 1-.706 0Z"/><path d="M20.04 6.505a.5.5 0 0 0-.583-.582l-2.6.485a.5.5 0 0 0-.262.138L10.2 12.964a.5.5 0 0 0 0 .707l2.126 2.117a.5.5 0 0 0 .707 0l6.394-6.419a.5.5 0 0 0 .138-.263l.474-2.601Zm-.766-1.565a1.5 1.5 0 0 1 1.75 1.744l-.474 2.602a1.5 1.5 0 0 1-.413.79l-6.394 6.417a1.5 1.5 0 0 1-2.122.004L9.496 14.38a1.5 1.5 0 0 1-.004-2.122l6.394-6.418a1.5 1.5 0 0 1 .788-.416l2.6-.484Z"/><path d="M6.31 12.618a2 2 0 0 1 2.83-.006l4.25 4.235a2 2 0 0 1-2.823 2.834l-4.251-4.235a2 2 0 0 1-.005-2.828Zm2.123.703a1 1 0 1 0-1.411 1.417l4.25 4.234a1 1 0 0 0 1.412-1.417l-4.25-4.234Z"/><path d="m8.121 16.533l-2.823 2.834a.5.5 0 0 0 .001.707l.709.706a.5.5 0 0 0 .707-.001l2.823-2.834l.709.706l-2.824 2.834a1.5 1.5 0 0 1-2.12.003l-.71-.705a1.5 1.5 0 0 1-.003-2.122l2.823-2.833l.708.705Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilSwordCircleFilled0)"/></g>`),
		g.Group(children),
	)
}