package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sleep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.953 14H1.027a.945.945 0 0 1-.867-.589a1.02 1.02 0 0 1 .168-1.063l5.531-6.375H.988C.465 5.973 0 5.572 0 5.028c0-.545.425-.985.947-.985h7.026c.374 0 .716.23.866.589c.15.359.085.774-.168 1.063l-5.492 6.387h4.773c.523 0 .947.389.947.933c.001.544-.422.985-.946.985zm7.802-5.065H9.701a.64.64 0 0 1-.586-.391a.675.675 0 0 1 .107-.711l5.116-5.869h-5c-.354 0-.318-.93.035-.93h6.098c.253 0 .446.117.55.354a.675.675 0 0 1-.107.712L10.88 7.963h4.875c.354.001.354.972 0 .972z"/>`),
		g.Group(children),
	)
}