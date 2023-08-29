package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Download(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.63 0 0 268.63 0 600s268.63 600 600 600c331.369 0 600-268.631 600-600C1200 268.63 931.369 0 600 0zm0 1069.565c-259.37 0-469.565-210.261-469.565-469.565S340.63 130.435 600 130.435c259.369 0 469.565 210.261 469.565 469.565S859.369 1069.565 600 1069.565zm117.392-720.652H482.608v266.739H335.87L600 864.13l264.13-248.478H717.391l.001-266.739z"/>`),
		g.Group(children),
	)
}