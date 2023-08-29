package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 8l-.002 2l-2.505 2.883a3.752 3.752 0 0 1-.993 7.367a3.751 3.751 0 0 1-3.682-3.033l1.964-.382a1.75 1.75 0 1 0 .924-1.895l-1.307-1.547L19.35 10H15V8h7ZM4 4v7h7V4h2v16h-2v-7H4v7H2V4h2Z"/>`),
		g.Group(children),
	)
}