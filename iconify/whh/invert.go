package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Invert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm0-768q0-53-38-91l-120 121q-44-45-102.5-69.5t-123.5-24.5q-87 0-160.5 43t-116.5 116.5t-43 160.5q0 65 24.5 123.5t69.5 102.5l-121 121q38 37 91 37h512q53 0 90.5-37.5t37.5-90.5V256zm-384 576q-133 0-226-94l452-452q45 44 69.5 102.5t24.5 123.5q0 87-43 160.5T672.928 789t-160.5 43z"/>`),
		g.Group(children),
	)
}