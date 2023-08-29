package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0zM250.049 310.474h699.902v462.013h-254.15v66.576h61.963v50.465H442.236v-50.465h61.963v-66.576h-254.15V310.474zm86.133 87.231v287.622h527.637V397.705H336.182z"/>`),
		g.Group(children),
	)
}