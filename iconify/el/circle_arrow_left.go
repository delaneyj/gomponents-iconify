package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 1200c331.359 0 600-268.641 600-600S931.359 0 600 0S0 268.641 0 600s268.641 600 600 600zm66.656-297.141L212.32 600l454.336-302.859v163.184H987.68v279.352H666.656v163.182z"/>`),
		g.Group(children),
	)
}