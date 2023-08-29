package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwordCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopSwordCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M19.418 4.209a2 2 0 0 1 2.333 2.325l-.474 2.601a2 2 0 0 1-.551 1.053l-6.394 6.418a2 2 0 0 1-2.828.005l-2.126-2.117a2 2 0 0 1-.005-2.829l6.394-6.418a2 2 0 0 1 1.05-.554l2.6-.484Zm-.109 4.567l.475-2.601l-2.6.484l-6.394 6.418l2.125 2.117l6.394-6.418Z"/><path d="M10.97 15.021a.75.75 0 0 1-.002-1.06l3.53-3.543a.75.75 0 1 1 1.062 1.059l-3.529 3.542a.75.75 0 0 1-1.06.002Z"/><path d="M6.192 12.025a2.5 2.5 0 0 1 3.535-.007l4.25 4.235a2.5 2.5 0 1 1-3.528 3.542l-4.25-4.235a2.5 2.5 0 0 1-.007-3.535Zm2.124 1.41a.5.5 0 0 0-.706.708l4.25 4.235a.5.5 0 1 0 .706-.708l-4.25-4.235Z"/><path d="M8.71 16.646L5.889 19.48l.708.706l2.823-2.834l1.417 1.412l-2.823 2.833a2 2 0 0 1-2.829.006l-.708-.706a2 2 0 0 1-.005-2.829l2.823-2.833l1.417 1.411Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopSwordCircleFilled0)"/></g>`),
		g.Group(children),
	)
}