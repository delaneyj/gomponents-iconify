package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.05 13h5.477a17.9 17.9 0 0 0 2.925 8.881a10.005 10.005 0 0 1-8.403-8.88Zm0-2a10.005 10.005 0 0 1 8.402-8.88A17.9 17.9 0 0 0 7.527 11H2.05Zm19.9 0h-5.477a17.9 17.9 0 0 0-2.925-8.88A10.005 10.005 0 0 1 21.95 11Zm0 2a10.005 10.005 0 0 1-8.402 8.881a17.9 17.9 0 0 0 2.925-8.88h5.478ZM9.53 13h4.94A15.908 15.908 0 0 1 12 20.592A15.908 15.908 0 0 1 9.53 13Zm0-2A15.908 15.908 0 0 1 12 3.41A15.908 15.908 0 0 1 14.47 11H9.53Z"/>`),
		g.Group(children),
	)
}