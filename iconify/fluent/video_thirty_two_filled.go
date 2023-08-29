package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20.5 10A4.5 4.5 0 0 0 16 5.5H6.5A4.5 4.5 0 0 0 2 10v12a4.5 4.5 0 0 0 4.5 4.5H16a4.5 4.5 0 0 0 4.5-4.5V10Zm1.5 9.8l5.013 3.76c1.23.923 2.987.045 2.987-1.493V9.934c0-1.539-1.756-2.417-2.987-1.494L22 12.2v7.6Z"/>`),
		g.Group(children),
	)
}