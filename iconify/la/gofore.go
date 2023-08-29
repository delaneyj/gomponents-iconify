package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gofore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 4C10.373 4 5 9.373 5 16s5.373 12 12 12c3.585 0 6.782-1.592 8.979-4.086c-.116-1.795-1.489-3.837-3.979-3.902v.853C20.729 22.175 18.964 23 17 23c-3.86 0-7-3.14-7-7s3.14-7 7-7a6.97 6.97 0 0 1 4.457 1.607l3.506-3.568A11.943 11.943 0 0 0 17 4zm-1 9v5h6c2.21 0 3.418.796 4 2c0-2.631.002-7-4-7h-6z"/>`),
		g.Group(children),
	)
}