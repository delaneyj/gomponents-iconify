package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M149 48V5h43v470h-43v-43H43q-18 0-30.5-12.5T0 389V91q0-18 12.5-30.5T43 48h106zm0 320V240L43 368h106zM341 48q18 0 30.5 12.5T384 91v298q0 18-12.5 30.5T341 432H235V240l106 128V91H235V48h106z"/>`),
		g.Group(children),
	)
}