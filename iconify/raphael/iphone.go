package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.755 1H10.62c-1.136 0-2.058.92-2.058 2.058v24.385c0 1.136.92 2.058 2.058 2.058h10.135c1.136 0 2.058-.92 2.058-2.057V3.058A2.058 2.058 0 0 0 20.755 1zM14.66 3.264h2.056c.1 0 .183.08.183.18c0 .1-.083.18-.184.18H14.66c-.1 0-.182-.08-.182-.18c0-.1.08-.18.18-.18zm-1.435-.206a.36.36 0 1 1 0 .72a.36.36 0 0 1 0-.72zm2.463 25.415a1.44 1.44 0 1 1 0-2.88a1.44 1.44 0 0 1 0 2.88zm6.353-4.118a.31.31 0 0 1-.308.31H9.642a.31.31 0 0 1-.308-.31V6.042c0-.17.138-.31.308-.31h12.09c.17 0 .31.14.31.31v18.313z"/>`),
		g.Group(children),
	)
}