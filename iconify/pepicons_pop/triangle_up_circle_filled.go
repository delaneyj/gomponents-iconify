package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleUpCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopTriangleUpCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M13 10.002L8.963 17h8.074L13 10.002Zm.866-2.5a1 1 0 0 0-1.732 0L6.365 17.5a1 1 0 0 0 .866 1.5H18.77a1 1 0 0 0 .866-1.5l-5.769-9.999Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopTriangleUpCircleFilled0)"/></g>`),
		g.Group(children),
	)
}