package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.43 5c3.551 0 6.43-1.242 6.43-2.579C13.86 1.084 10.981 0 7.43 0C3.879 0 1 1.084 1 2.421C1 3.758 3.879 5 7.43 5zm6.449 7.917a.716.716 0 0 0 .031-.204V9.444c0 .655-2.932 1.636-6.447 1.636s-6.447-.98-6.447-1.636v3.269c0 1.188 2.887 2.146 6.447 2.146c.885 0 1.728-.06 2.495-.166v-1.776h3.921z"/><path d="M7.494 9.919c3.561 0 6.447-.982 6.447-2.196V4.377c0 .672-2.932 1.674-6.447 1.674S1.047 5.049 1.047 4.377v3.346c0 1.214 2.887 2.196 6.447 2.196zM11 14h4v.912h-4z"/></g>`),
		g.Group(children),
	)
}