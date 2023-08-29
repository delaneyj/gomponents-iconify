package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Siren(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M120 16V8a8 8 0 0 1 16 0v8a8 8 0 0 1-16 0Zm80 32a8 8 0 0 0 5.66-2.34l8-8a8 8 0 0 0-11.32-11.32l-8 8A8 8 0 0 0 200 48ZM50.34 45.66a8 8 0 0 0 11.32-11.32l-8-8a8 8 0 0 0-11.32 11.32Zm87 26.45a8 8 0 1 0-2.64 15.78C153.67 91.08 168 108.32 168 128a8 8 0 0 0 16 0c0-27.4-20.07-51.43-46.68-55.89ZM232 176v24a16 16 0 0 1-16 16H40a16 16 0 0 1-16-16v-24a16 16 0 0 1 16-16v-32a88 88 0 0 1 88.67-88c48.15.36 87.33 40.29 87.33 89v31a16 16 0 0 1 16 16ZM56 160h144v-31c0-40-32.05-72.71-71.45-73H128a72 72 0 0 0-72 72Zm160 40v-24H40v24h176Z"/>`),
		g.Group(children),
	)
}