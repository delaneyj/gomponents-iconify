package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Html(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#cfa967"/><path fill="#fff" fill-rule="nonzero" d="M16.02 16.945v7.993l5.947-1.602l1.397-15.39H16.02v6.532l.304-.947h.796zm-7.384 7.524L7 6.5h18l-1.636 17.969L15.98 26.5zm5.163-6.793v-.741l-2.469-.984l2.47-.99v-.742l-3.527 1.433v.592zm7.933-1.432l-3.527 1.432v-.741l2.47-.987l-2.47-.987v-.742l3.527 1.433zm-5.712.7v-2.466l-1.317 4.107h.788z"/></g>`),
		g.Group(children),
	)
}