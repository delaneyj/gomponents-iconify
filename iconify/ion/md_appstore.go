package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdAppstore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M464 152H360c0-57-46.562-103.859-104-103.859S152 95 152 152H48c18.688 216 13 312 13 312h389.999c-.001 0-5.688-98 13.001-312zM256 74.105c43.008 0 77.999 34.895 77.999 77.895H178c0-43 34.991-77.895 78-77.895zM204 397.64V228.867l142.999 84.387L204 397.64z" fill="currentColor"/>`),
		g.Group(children),
	)
}