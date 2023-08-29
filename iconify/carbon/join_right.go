package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JoinRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 6a9.95 9.95 0 0 0-4.003.84a9.99 9.99 0 0 1 0 18.32A9.999 9.999 0 1 0 20 6Z"/><path fill="none" d="M20 16a8.01 8.01 0 0 0-6.248-7.8a9.976 9.976 0 0 1 2.245-1.36a9.99 9.99 0 0 1 0 18.32a9.976 9.976 0 0 1-2.245-1.36A8.01 8.01 0 0 0 20 16Z"/><path fill="currentColor" d="M10 16a9.976 9.976 0 0 0 3.752 7.8a7.994 7.994 0 0 0 0-15.6A9.976 9.976 0 0 0 10 16Z"/><path fill="currentColor" d="M12 24a8 8 0 1 1 1.752-15.8a9.976 9.976 0 0 1 2.245-1.36a10 10 0 1 0 0 18.32a9.976 9.976 0 0 1-2.245-1.36A7.973 7.973 0 0 1 12 24Z"/>`),
		g.Group(children),
	)
}