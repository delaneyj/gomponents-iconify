package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaxiBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M244 112h-9.45l-30.09-46.82A19.93 19.93 0 0 0 187.63 56h-19.51l-11-27.43A19.9 19.9 0 0 0 138.58 16h-21.16a19.9 19.9 0 0 0-18.57 12.57L87.88 56H68.37a19.91 19.91 0 0 0-16.83 9.19L21.45 112H12a12 12 0 0 0 0 24h4v68a20 20 0 0 0 20 20h24a20 20 0 0 0 20-20v-20h96v20a20 20 0 0 0 20 20h24a20 20 0 0 0 20-20v-68h4a12 12 0 0 0 0-24ZM120.12 40h15.76l6.4 16h-28.56ZM70.55 80h114.9L206 112H50ZM56 200H40v-16h16Zm144 0v-16h16v16Zm16-40H40v-24h176Z"/>`),
		g.Group(children),
	)
}