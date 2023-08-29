package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 1200C268.63 1200 0 931.369 0 600C0 268.63 268.63 0 600 0c331.369 0 600 268.63 600 600s-268.631 600-600 600zm0-1069.565c-259.369 0-469.565 210.261-469.565 469.565c0 259.305 210.196 469.564 469.565 469.564S1069.564 859.305 1069.564 600c0-259.304-210.195-469.565-469.564-469.565zm117.392 720.652H482.608V584.348H335.87L600 335.87l264.131 248.478H717.392v266.739z"/>`),
		g.Group(children),
	)
}