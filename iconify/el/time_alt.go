package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M1024.236 175.764C915.665 67.192 765.675 0 600 0C268.651 0 0 268.651 0 600c0 331.35 268.651 600 600 600c331.35 0 600-268.651 600-600c0-165.675-67.192-315.665-175.764-424.236zm-121.75 86.328l108.13 108.131l-405.851 405.851l-.052-.053l-.052.053l-284.153-284.153l11.859-11.858l84.05-84.05l12.222-12.221l176.074 176.074l297.773-297.774z"/>`),
		g.Group(children),
	)
}