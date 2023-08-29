package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DkOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#c8102e" d="M0 0h512.1v512H0z"/><path fill="#fff" d="M144 0h73.1v512H144z"/><path fill="#fff" d="M0 219.4h512.1v73.2H0z"/>`),
		g.Group(children),
	)
}