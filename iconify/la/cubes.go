package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cubes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 4l-.375.156l-6 2.407l-.625.25v6.718l-5.406 2.344l-.594.281v8.063l.5.312l6 3.344l.469.25l.469-.219L16 25.125l5.563 2.781l.468.219l.469-.25l6-3.344l.5-.312v-8.063l-.594-.281L23 13.531V6.812l-.625-.25l-6-2.406zm0 2.188l3.281 1.28L16 8.75l-3.281-1.281zm-5 2.75l4 1.53v4.876l-4-1.781zm10 0v4.624l-4 1.782v-4.875zm-11 6.374l3.625 1.594L10 18.688l-3.625-1.813zm12 0l3.625 1.563L22 18.688l-3.625-1.782l1.125-.5zM5 18.407l4 2v4.907l-4-2.25zm22 0v4.657l-4 2.25v-4.907zm-12 .063v4.906l-4 2v-4.969zm2 0l4 1.937v4.969l-4-2z"/>`),
		g.Group(children),
	)
}