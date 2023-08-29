package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M15.752 24a3.25 3.25 0 1 1-6.5 0a3.25 3.25 0 0 1 6.5 0z" fill="currentColor"/><path d="M27.252 24a3.25 3.25 0 1 1-6.5 0a3.25 3.25 0 0 1 6.5 0z" fill="currentColor"/><path d="M35.502 27.25a3.25 3.25 0 1 0 0-6.5a3.25 3.25 0 0 0 0 6.5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}