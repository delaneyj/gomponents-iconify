package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyringeCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilSyringeCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M6.207 14.571a2 2 0 0 0 0 2.829l1.414 1.414a2 2 0 0 0 2.829 0l6.97-6.971a2 2 0 0 0 .586-1.46l-.032-1.407a2 2 0 0 0-1.978-1.955l-1.375-.014a2 2 0 0 0-1.435.585l-6.979 6.98Zm10.507-3.436l-6.971 6.972a1 1 0 0 1-1.415 0l-1.414-1.415a1 1 0 0 1 0-1.414L13.893 8.3a1 1 0 0 1 .718-.292l1.374.014a1 1 0 0 1 .99.978l.031 1.407a1 1 0 0 1-.292.73Z" clip-rule="evenodd"/><path d="M12.52 15.107a.5.5 0 1 1-.706.707l-1.415-1.415a.5.5 0 1 1 .708-.707l1.414 1.415Zm-5.956 5.657a.5.5 0 0 1-.707.707l-2.529-2.529a.5.5 0 1 1 .708-.707l2.528 2.529Zm3.79-.118a.5.5 0 1 1-.708.707l-6-6a.5.5 0 1 1 .708-.707l6 6Zm4.288-7.661a.5.5 0 0 1-.707.707l-1.414-1.414a.5.5 0 0 1 .707-.707l1.414 1.414Z"/><path d="m5 20.457l-.707-.707l2.457-2.457l.707.707L5 20.457ZM19.854 4.354a.5.5 0 1 1 .707.707l-3 3a.5.5 0 1 1-.707-.707l3-3Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilSyringeCircleFilled0)"/></g>`),
		g.Group(children),
	)
}