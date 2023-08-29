package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adjust(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.641 0 0 268.641 0 600s268.641 600 600 600s600-268.641 600-600S931.359 0 600 0zm0 174.975V1025.1c-234.75 0-425.109-190.336-425.109-425.109c0-234.751 190.336-425.016 425.109-425.016z"/>`),
		g.Group(children),
	)
}