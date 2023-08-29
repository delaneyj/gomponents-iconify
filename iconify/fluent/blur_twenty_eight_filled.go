package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlurTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 14C3 7.925 7.925 3 14 3c2.024 0 3.92.547 5.549 1.5H14v1h6.983c.547.45 1.051.953 1.503 1.5H14v1h9.221c.31.476.585.977.82 1.5H14v1h10.431c.163.486.293.987.388 1.5H14v1h10.955a11.17 11.17 0 0 1 .034 1.5H14v1h10.899c-.07.512-.175 1.013-.313 1.5H14v1h10.25c-.203.52-.445 1.022-.722 1.5H14v1h8.875c-.396.54-.84 1.042-1.325 1.5H14v1h6.326A10.95 10.95 0 0 1 14 25C7.925 25 3 20.075 3 14Z"/>`),
		g.Group(children),
	)
}