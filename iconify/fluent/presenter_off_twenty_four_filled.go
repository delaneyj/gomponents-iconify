package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenterOffTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.28 2.22a.75.75 0 1 0-1.06 1.06l5.854 5.855A1.75 1.75 0 0 0 7 10.75V12h3.94l.996.997H3.75c-.709 0-1.022.893-.469 1.336L8 18.11v1.645a2.25 2.25 0 0 0 2.25 2.25h3.495a2.25 2.25 0 0 0 2.25-2.25V18.11l.586-.468l4.138 4.139a.75.75 0 0 0 1.061-1.061L3.28 2.22Zm12.9 10.777l2.758 2.758l1.776-1.422c.553-.444.24-1.336-.469-1.336H16.18ZM12.182 9l3 3h1.813v-1.25l-.006-.143A1.75 1.75 0 0 0 15.245 9h-3.063ZM12 8c-.344 0-.674-.057-.982-.164L9.164 5.982A3 3 0 1 1 12 8Z"/>`),
		g.Group(children),
	)
}