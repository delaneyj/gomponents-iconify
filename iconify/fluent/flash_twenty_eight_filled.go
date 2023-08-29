package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.406 2.852A1.333 1.333 0 0 1 9.65 2h8.516c.927 0 1.57.922 1.252 1.792L17.324 9.5h4.837c1.178 0 1.777 1.416.957 2.262L9.784 25.503c-1.14 1.175-3.106.117-2.753-1.482l1.66-7.521H5.917a1.917 1.917 0 0 1-1.787-2.61L8.406 2.853Z"/>`),
		g.Group(children),
	)
}