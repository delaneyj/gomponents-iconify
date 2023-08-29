package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fontcomic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M864 256H569q-10 88-33.5 224t-51 282T448 960q-33 17-101 40.5t-91 23.5q2-15 15-90.5t15.5-92t13-80.5t13.5-87t10.5-77.5t10-85.5t7-78.5t6.5-89t3-86.5q-57 4-137 20T83 306l-51 14L0 64q172-16 384-32q86-6 214-14T811 5l85-5z"/>`),
		g.Group(children),
	)
}