package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SexMale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M512 512q-212 0-362 150T0 1024t150 362t362 150t362-150t150-362q0-140-71-260l327-327v395h256V0H704v257h394L772 583q-120-71-260-71zm0 256q106 0 181 75t75 181t-75 181t-181 75t-181-75t-75-181t75-181t181-75z"/>`),
		g.Group(children),
	)
}