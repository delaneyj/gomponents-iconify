package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListLtrNinetyTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 3.25a1.25 1.25 0 1 1-2.498 0a1.25 1.25 0 0 1 2.499 0ZM6 6.75v14.5a.75.75 0 0 1-1.493.102l-.006-.102V6.75a.75.75 0 0 1 1.493-.102L6 6.75Zm7-3.5a1.25 1.25 0 1 1-2.498 0a1.25 1.25 0 0 1 2.498 0Zm-.5 3.5v14.5a.75.75 0 0 1-1.493.102l-.006-.102V6.75a.75.75 0 0 1 1.493-.102l.007.102Zm7-3.5a1.25 1.25 0 1 1-2.498 0a1.25 1.25 0 0 1 2.498 0Zm-.5 3.5v14.5a.75.75 0 0 1-1.493.102l-.007-.102V6.75a.75.75 0 0 1 1.494-.102L19 6.75Z"/>`),
		g.Group(children),
	)
}