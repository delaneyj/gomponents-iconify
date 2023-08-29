package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GpsF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.919 2.635l-5.953 16.08c-.376 1.016-1.459 1.538-2.418 1.165a1.851 1.851 0 0 1-1.045-1.054l-1.887-4.77a3.712 3.712 0 0 0-1.955-2.052l-4.542-1.981C.174 9.61-.256 8.465.157 7.465a1.97 1.97 0 0 1 1.067-1.079L16.54.136c.967-.395 2.04.101 2.395 1.109c.157.446.151.94-.015 1.39z"/>`),
		g.Group(children),
	)
}