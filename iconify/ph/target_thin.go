package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TargetThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M218.26 84.89a100.16 100.16 0 1 1-22.44-30.37l25.35-25.35a4 4 0 1 1 5.66 5.66l-96 96a4 4 0 0 1-5.66-5.66l31-31a44 44 0 1 0 15.78 31.3a4 4 0 0 1 8-.46a52 52 0 1 1-18.1-36.51l28.34-28.33A92 92 0 0 0 63 193.05A92 92 0 0 0 211 88.33a4 4 0 1 1 7.22-3.44Z"/>`),
		g.Group(children),
	)
}