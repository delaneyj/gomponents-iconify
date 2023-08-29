package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallpaper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 45v150H0V45q0-17 12.5-29.5T43 3h149v42H43zm128 192l63 79l43-57l64 85H85zm149-95.5q0 13.5-9.5 22.5t-22.5 9t-22.5-9t-9.5-22.5t9.5-23T288 109t22.5 9.5t9.5 23zM384 3q18 0 30.5 12.5T427 45v150h-43V45H235V3h149zm0 384V237h43v150q0 17-12.5 29.5T384 429H235v-42h149zM43 237v150h149v42H43q-18 0-30.5-12.5T0 387V237h43z"/>`),
		g.Group(children),
	)
}