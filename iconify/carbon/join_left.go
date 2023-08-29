package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JoinLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 6a9.927 9.927 0 0 0-3.997.84a9.976 9.976 0 0 1 2.245 1.36a8 8 0 1 1 0 15.601a9.976 9.976 0 0 1-2.245 1.36A9.998 9.998 0 1 0 20 6Z"/><path fill="currentColor" d="M12 16a8.01 8.01 0 0 0 6.248 7.8a9.986 9.986 0 0 0 0-15.6A8.01 8.01 0 0 0 12 16Z"/><path fill="none" d="M12 16a8.01 8.01 0 0 1 6.248-7.8a9.976 9.976 0 0 0-2.245-1.36a9.99 9.99 0 0 0 0 18.32a9.976 9.976 0 0 0 2.245-1.36A8.01 8.01 0 0 1 12 16Z"/><path fill="currentColor" d="M10 16a10.01 10.01 0 0 1 6.003-9.16a10 10 0 1 0 0 18.32A10.01 10.01 0 0 1 10 16Z"/>`),
		g.Group(children),
	)
}