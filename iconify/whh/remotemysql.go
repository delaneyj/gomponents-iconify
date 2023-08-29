package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Remotemysql(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 132v605l64 64q-145 95-318 95q-137 0-261-63l-6 12q73 74 73 179H0q0-106 75-181t181-75q30 0 62 8l11-21q-94-81-147.5-194.5T128 318q0-173 95-318l65 64h602l49-50q15-14 35-14t34.5 14.5T1023 49t-14 35zm-609-4l238 237l237-237H351zm545 68L658 435l238 238V196z"/>`),
		g.Group(children),
	)
}