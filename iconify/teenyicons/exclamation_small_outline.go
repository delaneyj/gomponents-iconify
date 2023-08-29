package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationSmallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M8 10.5V10H7v.5h1Zm-1 .01v.5h1v-.5H7ZM7 4v4h1V4H7Zm0 6.5v.01h1v-.01H7Z"/>`),
		g.Group(children),
	)
}