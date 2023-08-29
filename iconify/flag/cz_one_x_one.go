package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CzOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fff" d="M0 0h512v256H0z"/><path fill="#d7141a" d="M0 256h512v256H0z"/><path fill="#11457e" d="M300 256L0 56v400z"/>`),
		g.Group(children),
	)
}