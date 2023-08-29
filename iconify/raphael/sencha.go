package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sencha(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18.265 22.734a4.09 4.09 0 0 1 .133 7.297l1.922-.98a9.63 9.63 0 0 0 5.332-8.616c0-3.72-2.11-6.945-5.195-8.547l-6.272-3.144a4.097 4.097 0 0 1-2.308-3.684c0-1.566.88-2.927 2.175-3.613l-1.922.98a9.63 9.63 0 0 0-.137 17.163l6.272 3.144z"/>`),
		g.Group(children),
	)
}