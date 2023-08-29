package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Roundcube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M736 1024L0 704V448l193-110q16-144 116-241T544 0q136 0 236.5 98T895 341l129 43v256zM544 64q-112 0-194.5 77.5T256 333q0 25 40 52t108 45t140 18t140-18t108-45t40-52q-11-114-93.5-191.5T544 64z"/>`),
		g.Group(children),
	)
}