package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CookieLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M163.07 164.93a10 10 0 1 1-14.14 0a10 10 0 0 1 14.14 0Zm-78.14-8a10 10 0 1 0 14.14 0a10 10 0 0 0-14.14 0Zm6.14-41.86a10 10 0 1 0-14.14 0a10 10 0 0 0 14.14 0Zm33.86 1.86a10 10 0 1 0 14.14 0a10 10 0 0 0-14.14 0ZM230 128A102 102 0 1 1 128 26a6 6 0 0 1 6 6a42 42 0 0 0 42 42a6 6 0 0 1 6 6a42 42 0 0 0 42 42a6 6 0 0 1 6 6Zm-12.18 5.65A54.09 54.09 0 0 1 170.3 85.7a54.09 54.09 0 0 1-48-47.53a90 90 0 1 0 95.47 95.48Z"/>`),
		g.Group(children),
	)
}