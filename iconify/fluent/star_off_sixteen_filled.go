package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOffSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m12.358 13.065l1.788 1.789a.5.5 0 0 0 .708-.708l-13-13a.5.5 0 1 0-.708.708l3.482 3.482l-2.356.342a.9.9 0 0 0-.5 1.535l2.462 2.4L3.653 13a.9.9 0 0 0 1.306.948L8 12.35l3.042 1.6a.9.9 0 0 0 1.315-.884Zm-.59-3.453l.007.042L6.212 4.09l.982-1.99a.9.9 0 0 1 1.614 0l1.521 3.083l3.401.494a.9.9 0 0 1 .5 1.535l-2.462 2.4Z"/>`),
		g.Group(children),
	)
}