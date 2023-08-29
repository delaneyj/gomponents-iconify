package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Percentage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.05 17.7a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06l9.9-9.9a.75.75 0 1 1 1.06 1.06l-9.9 9.9a.74.74 0 0 1-.53.22Zm1.45-6.95a2.25 2.25 0 1 1 2.25-2.25a2.25 2.25 0 0 1-2.25 2.25Zm0-3a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75Zm7 10a2.25 2.25 0 1 1 2.25-2.25a2.25 2.25 0 0 1-2.25 2.25Zm0-3a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75Z"/>`),
		g.Group(children),
	)
}