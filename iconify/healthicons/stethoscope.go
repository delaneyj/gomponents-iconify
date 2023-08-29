package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stethoscope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M12 10a2 2 0 1 0-1.732-3H10a3 3 0 0 0-3 3v9c-.552 0-1.005.449-.955.999A11 11 0 0 0 14 29.583V32c0 1.306.835 2.418 2 2.83a7.25 7.25 0 0 0 14.5-.08v-4a3.75 3.75 0 1 1 7.5 0v2.42a3.001 3.001 0 1 0 2 0v-2.42a5.75 5.75 0 0 0-11.5 0v4a5.25 5.25 0 0 1-10.5.08A3.001 3.001 0 0 0 20 32v-2.417a11 11 0 0 0 7.955-9.584c.05-.55-.403-.999-.955-.999v-9a3 3 0 0 0-3-3h-.268A2 2 0 0 0 20 8a2 2 0 0 0 3.732 1H24a1 1 0 0 1 1 1v9h.21c-.553 0-.993.45-1.07.997a7.21 7.21 0 0 1-14.28 0C9.783 19.45 9.343 19 8.79 19H9v-9a1 1 0 0 1 1-1h.268A2 2 0 0 0 12 10Z"/>`),
		g.Group(children),
	)
}