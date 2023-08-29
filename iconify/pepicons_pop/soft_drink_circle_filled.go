package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoftDrinkCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopSoftDrinkCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M16.411 7H9.589a2.502 2.502 0 0 0-2.481 2.803l1.22 10A2.5 2.5 0 0 0 10.809 22h4.385a2.5 2.5 0 0 0 2.481-2.198l1.218-10A2.5 2.5 0 0 0 16.41 7ZM9.529 9.004A.5.5 0 0 1 9.589 9h6.822a.5.5 0 0 1 .496.56l-1.217 10a.5.5 0 0 1-.496.44h-4.385a.5.5 0 0 1-.496-.44l-1.22-10a.5.5 0 0 1 .436-.556Z" clip-rule="evenodd"/><path d="M12.217 17.72a.75.75 0 0 1-1.434-.44l4-13a.75.75 0 1 1 1.434.44l-4 13Z"/><path d="M8.5 13.25a.75.75 0 0 1 0-1.5h9a.75.75 0 0 1 0 1.5h-9Zm6.818-8.022a.75.75 0 0 1 .364-1.456l4 1a.75.75 0 1 1-.364 1.456l-4-1Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopSoftDrinkCircleFilled0)"/></g>`),
		g.Group(children),
	)
}