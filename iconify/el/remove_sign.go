package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0zM411.475 262.5L600 451.025L788.525 262.5L937.5 411.475L748.975 600L937.5 788.525L788.525 937.5L600 748.975L411.475 937.5L262.5 788.525L451.025 600L262.5 411.475L411.475 262.5z"/>`),
		g.Group(children),
	)
}