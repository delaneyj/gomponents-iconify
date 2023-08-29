package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M28.034 17.29c.13.43.53.71.96.71v-.01a.993.993 0 0 0 .96-1.28C29.923 16.61 26.613 6 15.995 6S2.07 16.61 2.04 16.72c-.16.53.14 1.08.67 1.24c.53.16 1.09-.14 1.25-.67C4.07 16.91 6.89 8 15.997 8c9.104 0 11.915 8.903 12.037 9.29ZM10 18a6 6 0 1 1 12 0a6 6 0 0 1-12 0Z"/>`),
		g.Group(children),
	)
}