package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForDjibouti(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#009e60" d="M32.454 32.03L4.11 52.25c1.538 1.106 3.499 1.754 5.89 1.754h44c6.627 0 10-4.925 10-11v-11H32.416l.038.027"/><path fill="#6ab2e7" d="M54 10H10c-2.201 0-4.04.553-5.514 1.5L32.416 32H64V21c0-6.075-3.373-11-10-11"/><path fill="#e6e7e8" d="M32.416 32L4.486 11.5C1.513 13.407 0 16.942 0 21v22c0 3.883 1.385 7.289 4.11 9.246l28.344-20.219l-.038-.027z"/><path fill="#be1e2d" d="m21.06 28.357l-6.303.016l-1.95-6.379l-1.945 6.379l-6.308-.016l5.109 3.883l-1.976 6.334l5.124-3.936l5.132 3.936l-1.982-6.334z"/>`),
		g.Group(children),
	)
}