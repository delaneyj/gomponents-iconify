package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyringeCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopSyringeCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M7.354 13.717a2.5 2.5 0 0 0 0 3.536l1.414 1.414a2.5 2.5 0 0 0 3.535 0l6.971-6.97a2.5 2.5 0 0 0 .732-1.825l-.032-1.407a2.5 2.5 0 0 0-2.473-2.444l-1.374-.014a2.5 2.5 0 0 0-1.795.732l-6.978 6.978Zm10.506-3.435l-6.97 6.971a.5.5 0 0 1-.708 0l-1.414-1.414a.5.5 0 0 1 0-.707l6.978-6.979a.5.5 0 0 1 .36-.146l1.374.014a.5.5 0 0 1 .495.489l.032 1.407a.5.5 0 0 1-.147.365Z" clip-rule="evenodd"/><path d="M14.02 14.607a.5.5 0 1 1-.706.707l-1.415-1.415a.5.5 0 1 1 .708-.707l1.414 1.415ZM8.718 19.91a1 1 0 1 1-1.415 1.414l-2.828-2.828a1 1 0 1 1 1.414-1.415l2.829 2.829Zm3.667-.24a1 1 0 1 1-1.414 1.415l-5.964-5.964a1 1 0 1 1 1.414-1.414l5.964 5.964Zm3.757-7.185a.5.5 0 1 1-.707.707l-1.414-1.414a.5.5 0 1 1 .707-.707l1.414 1.414Z"/><path d="M8 19.414L6.586 18L8 16.586L9.414 18L8 19.414ZM20.97 3.97a.75.75 0 0 1 1.061 1.061l-3 3a.75.75 0 0 1-1.06-1.06l3-3Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopSyringeCircleFilled0)"/></g>`),
		g.Group(children),
	)
}