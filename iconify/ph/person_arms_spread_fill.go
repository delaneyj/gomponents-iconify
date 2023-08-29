package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonArmsSpreadFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M100 36a28 28 0 1 1 28 28a28 28 0 0 1-28-28Zm127.6 56.57A15.7 15.7 0 0 0 212 80H44a16 16 0 0 0-6.7 30.53h.06l53.89 23.73l-21.92 83.3a16 16 0 0 0 7.9 20.91A15.83 15.83 0 0 0 84 240a16 16 0 0 0 14.44-9.06L128 180l29.58 51a16 16 0 0 0 29.07-13.35l-21.92-83.3l54-23.76a15.7 15.7 0 0 0 8.87-18.02Z"/>`),
		g.Group(children),
	)
}