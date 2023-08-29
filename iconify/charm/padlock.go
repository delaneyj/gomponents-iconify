package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Padlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 6.75h10.5v7.5H2.75zm2-.5s-1-4.5 3.25-4.5s3.25 4.5 3.25 4.5"/>`),
		g.Group(children),
	)
}