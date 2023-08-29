package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M0 0h512v512H0Z"/><path fill="#fff" d="M0 0h512v341.3H0Z"/><path fill="#ce1126" d="M0 0h512v170.7H0Z"/><path fill="#007a3d" d="M86.4 320L128 192l41.6 128l-108.9-79.1h134.6M342.4 320L384 192l41.6 128l-108.9-79.1h134.6"/>`),
		g.Group(children),
	)
}