package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EsCtOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fcdd09" d="M0 0h512v512H0z"/><path stroke="#da121a" stroke-width="60" d="M0 90h810m0 120H0m0 120h810m0 120H0" transform="scale(.6321 .94815)"/>`),
		g.Group(children),
	)
}