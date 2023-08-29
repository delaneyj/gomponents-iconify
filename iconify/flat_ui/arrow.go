package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<clipPath id="flatUiArrow0"><circle cx="50" cy="50" r="50"/></clipPath><g fill-rule="evenodd" clip-path="url(#flatUiArrow0)" clip-rule="evenodd"><circle cx="50" cy="50" r="50" fill="#24AE5F"/><path fill="#E57D25" d="M34 50.5a2.5 2.5 0 0 0 2.5 2.5h66a2.5 2.5 0 1 0 0-5h-66a2.5 2.5 0 0 0-2.5 2.5z"/><path fill="#FBF063" d="m23 31l16 17h44L67 31H23z"/><path fill="#F0C419" d="m67 70l16-17H39L23 70h44z"/></g>`),
		g.Group(children),
	)
}