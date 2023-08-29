package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Error(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0zm0 164.062c240.762 0 435.938 195.176 435.938 435.938S840.762 1035.938 600 1035.938S164.062 840.762 164.062 600S359.238 164.062 600 164.062zM281.47 482.153v235.693h637.06V482.153H281.47z"/>`),
		g.Group(children),
	)
}