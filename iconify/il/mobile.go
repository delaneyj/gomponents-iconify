package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M0 49q0-19 14-33T46 3h325q19 0 33 13t13 33v579q0 20-13 33t-33 14H46q-19 0-32-14T0 628V49zm208 579q20 0 34-13t13.5-32.5T242 549t-34-13q-19 0-32 13t-13.5 33t13.5 33t32 13zm128-139q11 0 11-11V84q0-5-3-9t-8-3H81q-5 0-8 3t-4 9v394q0 5 4 8t8 3h255z"/>`),
		g.Group(children),
	)
}