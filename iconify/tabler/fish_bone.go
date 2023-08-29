package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FishBone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.69 7.44A6.973 6.973 0 0 0 15 12a6.97 6.97 0 0 0 1.699 4.571c1.914-.684 3.691-2.183 5.301-4.565c-1.613-2.384-3.394-3.883-5.312-4.565M2 9.504a40.73 40.73 0 0 0 2.422 2.504A39.679 39.679 0 0 0 2 14.506M18 11v.01M4.422 12H15m-8-2v4m4-6v8"/>`),
		g.Group(children),
	)
}