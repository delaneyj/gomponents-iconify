package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RubyTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5.873 3.26A.748.748 0 0 1 6.44 3h11.31c.223 0 .434.099.576.27l5 6a.75.75 0 0 1-.028.992l-10.75 11.5a.75.75 0 0 1-1.096 0l-10.75-11.5a.75.75 0 0 1-.02-1.003l5.19-6Zm.91 1.24L2.258 9.73L12 20.153l9.75-10.43L17.399 4.5Z"/>`),
		g.Group(children),
	)
}