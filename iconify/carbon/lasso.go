package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lasso(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 2h-8a9.984 9.984 0 0 0-4.965 18.655C7.025 20.77 7 20.882 7 21a3.993 3.993 0 0 0 2.91 3.83A4.005 4.005 0 0 1 6 28H4v2h2a6.004 6.004 0 0 0 5.928-5.12a3.997 3.997 0 0 0 2.93-2.88H20a10 10 0 0 0 0-20Zm-9 21a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Zm9-3h-5.142a3.984 3.984 0 0 0-7.15-1.264A7.99 7.99 0 0 1 12 4h8a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}