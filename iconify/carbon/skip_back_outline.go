package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipBackOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 22a1.004 1.004 0 0 1-.486-.126l-9-5a1 1 0 0 1 0-1.748l9-5A1 1 0 0 1 24 11v10a1 1 0 0 1-1 1Zm-6.94-6L22 19.3v-6.6ZM11 22H9V10h2z"/><path fill="currentColor" d="M16 30a14 14 0 1 1 14-14a14.016 14.016 0 0 1-14 14Zm0-26a12 12 0 1 0 12 12A12.014 12.014 0 0 0 16 4Z"/>`),
		g.Group(children),
	)
}