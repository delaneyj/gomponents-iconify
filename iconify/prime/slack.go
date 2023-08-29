package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.36 14.11a1.68 1.68 0 1 1-1.68-1.68h1.68Zm.85 0a1.68 1.68 0 0 1 3.36 0v4.21a1.68 1.68 0 0 1-3.36 0Zm1.68-6.75a1.68 1.68 0 1 1 1.68-1.68v1.68Zm0 .85a1.68 1.68 0 0 1 0 3.36H5.68a1.68 1.68 0 0 1 0-3.36Zm6.75 1.68a1.68 1.68 0 1 1 1.68 1.68h-1.68V9.89Zm-.85 0a1.68 1.68 0 0 1-3.36 0V5.68a1.68 1.68 0 0 1 3.36 0v4.21Zm-1.68 6.75a1.68 1.68 0 1 1-1.68 1.68v-1.68Zm0-.85a1.68 1.68 0 0 1 0-3.36h4.21a1.68 1.68 0 0 1 0 3.36Z"/>`),
		g.Group(children),
	)
}