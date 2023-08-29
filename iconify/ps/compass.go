package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0zm0 469q-88 0-150.5-62.5T43 256t62.5-150.5T256 43t150.5 62.5T469 256t-62.5 150.5T256 469zm-85-245v203l162-122q8-5 8-17V85L179 207q-8 5-8 17zm42 19q0-9 7-13l91-85l-12 124q0 9-7 13l-91 85zm64 13q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15z"/>`),
		g.Group(children),
	)
}