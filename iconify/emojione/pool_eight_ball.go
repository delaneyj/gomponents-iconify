package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoolEightBall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#333"/><ellipse cx="36.3" cy="24.3" fill="#f5f5f5" rx="13.1" ry="13.9" transform="rotate(-39.592 36.265 24.268)"/><path fill="#3e4347" d="M45.3 23.2c1.8 2.9.8 6.3-2.9 8.6s-7.2 1.7-9-1.2c-1.1-1.8-1.1-3.8 0-5.6c-1.7 0-3.1-.7-4.1-2.2c-1.7-2.7-.7-6 2.5-7.9c3.1-1.9 6.5-1.4 8.2 1.3c1 1.6 1 3.2.2 4.6c2.1-.3 3.9.6 5.1 2.4m-3.1 1.6c-.9-1.5-2.9-1.9-4.6-.8c-1.8 1.1-2.3 3-1.4 4.5c1 1.6 2.9 2 4.7.9c1.7-1.1 2.2-3.1 1.3-4.6M32 21.3c.9 1.4 2.7 1.9 4.3.8c1.6-1 2-2.8 1.1-4.2c-.9-1.4-2.7-1.8-4.3-.8c-1.5 1-1.9 2.7-1.1 4.2"/>`),
		g.Group(children),
	)
}