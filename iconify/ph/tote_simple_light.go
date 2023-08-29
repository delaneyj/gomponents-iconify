package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToteSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M234.47 70.73A14.09 14.09 0 0 0 223.92 66H174v-2a46 46 0 0 0-92 0v2H32.08a14 14 0 0 0-14 15.64l14.25 120a14.06 14.06 0 0 0 14 12.36h163.34a14.06 14.06 0 0 0 14-12.36l14.25-120a14 14 0 0 0-3.45-10.91ZM94 64a34 34 0 0 1 68 0v2H94Zm117.73 136.23a2 2 0 0 1-2.06 1.77H46.33a2 2 0 0 1-2.06-1.77L30 80.23a1.92 1.92 0 0 1 .49-1.53a2.07 2.07 0 0 1 1.58-.7h191.85a2.07 2.07 0 0 1 1.58.7a1.92 1.92 0 0 1 .49 1.53Z"/>`),
		g.Group(children),
	)
}