package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00A6ED" d="M2 6a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6Z"/><path fill="#fff" d="m8.707 10.293l1.586-1.586a1 1 0 0 1 1.414 0l7.586 7.586l.018.018c.123.13.345.176.472.05l1.493-1.494c.628-.628 1.702-.186 1.707.702l.033 6.441a1 1 0 0 1-1.005 1.006l-6.442-.033c-.888-.005-1.33-1.08-.702-1.707l1.493-1.493c.127-.127.081-.349-.049-.472a1.545 1.545 0 0 1-.018-.018l-7.586-7.586a1 1 0 0 1 0-1.414Z"/></g>`),
		g.Group(children),
	)
}