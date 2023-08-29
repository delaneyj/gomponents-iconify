package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Font(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M925 974H661q-15-40-18-92q-120 119-306 119q-95 0-169-29T47 877T0 715q0-40 8.5-74t21-59.5T65 534t43-36.5t53.5-27.5t57.5-20.5t64-15t63.5-11T412 415q50-6 100-15q24-4 37-7t34-11t32-17t19.5-25.5T643 301q0-126-170-126q-87 0-130.5 31.5T286 320H42q6-90 44-154t100.5-99T321 16.5T480 1q68 0 124.5 7.5T716 37t92.5 54.5t60.5 89T892 309v408q0 64 .5 91t8.5 77.5t24 88.5zM641 620V514q-33 19-87 30.5T453.5 563T363 582.5T293 625t-26 79q0 62 40 94.5T410 831q103 0 167-55t64-156z"/>`),
		g.Group(children),
	)
}