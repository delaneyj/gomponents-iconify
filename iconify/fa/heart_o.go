package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1664 468q0-81-21.5-143t-55-98.5T1506 167t-94-31t-98-8t-112 25.5t-110.5 64t-86.5 72t-60 61.5q-18 22-49 22t-49-22q-24-28-60-61.5t-86.5-72t-110.5-64T478 128t-98 8t-94 31t-81.5 59.5t-55 98.5T128 468q0 168 187 355l581 560l580-559q188-188 188-356zm128 0q0 221-229 450l-623 600q-18 18-44 18t-44-18L228 916q-10-8-27.5-26T145 824.5T77 727T23.5 606T0 468q0-220 127-344T478 0q62 0 126.5 21.5t120 58T820 148t76 68q36-36 76-68t95.5-68.5t120-58T1314 0q224 0 351 124t127 344z"/>`),
		g.Group(children),
	)
}