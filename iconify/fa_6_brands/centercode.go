package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Centercode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M329.2 268.6c-3.8 35.2-35.4 60.6-70.6 56.8c-35.2-3.8-60.6-35.4-56.8-70.6c3.8-35.2 35.4-60.6 70.6-56.8c35.1 3.8 60.6 35.4 56.8 70.6zm-85.8 235.1C96.7 496-8.2 365.5 10.1 224.3c11.2-86.6 65.8-156.9 139.1-192c161-77.1 349.7 37.4 354.7 216.6c4.1 147-118.4 262.2-260.5 254.8zm179.9-180c27.9-118-160.5-205.9-237.2-234.2c-57.5 56.3-69.1 188.6-33.8 344.4c68.8 15.8 169.1-26.4 271-110.2z"/>`),
		g.Group(children),
	)
}