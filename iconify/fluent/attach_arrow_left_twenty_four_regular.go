package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttachArrowLeftTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M11.772 3.743a6 6 0 0 1 8.867 8.063a6.461 6.461 0 0 0-1.486-.595l.044-.043a4.502 4.502 0 0 0-6.198-6.524l-.168.16l-.013.014l-9.536 9.536a.75.75 0 0 1-1.133-.977l.072-.084l9.549-9.55h.002z" fill="currentColor"/><path d="M11.212 19.15c.14.535.345 1.042.608 1.513l-.377.377l-.037.03a3.723 3.723 0 0 1-5.489-4.973a.764.764 0 0 1 .085-.13l.054-.06l.086-.088l.142-.148l.002.003l7.436-7.454a.75.75 0 0 1 .977-.074l.084.073a.75.75 0 0 1 .074.976l-.073.084l-7.594 7.613a2.23 2.23 0 0 0 3.174 3.106l.848-.848z" fill="currentColor"/><path d="M23 17.5a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0zm-8.5-.5a.5.5 0 0 0 0 1h4.793l-1.646 1.646a.5.5 0 1 0 .707.707l2.5-2.5a.5.5 0 0 0 0-.707l-2.5-2.5a.5.5 0 1 0-.707.707L19.293 17H14.5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}