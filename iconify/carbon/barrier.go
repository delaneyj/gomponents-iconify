package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barrier(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<defs/><path d="M15 5h2v4h-2z" fill="currentColor"/><path d="M15 11h2v4h-2z" fill="currentColor"/><path d="M15 17h2v4h-2z" fill="currentColor"/><path d="M15 23h2v4h-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}