package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockAlarmThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 3c1.523 0 2.886.68 3.803 1.754A11.975 11.975 0 0 1 16 4c1.476 0 2.89.267 4.197.754a5 5 0 1 1 7.05 7.05C27.733 13.11 28 14.524 28 16c0 2.954-1.067 5.658-2.837 7.749l2.544 2.544a1 1 0 0 1-1.414 1.414l-2.544-2.544A11.952 11.952 0 0 1 16 28a11.953 11.953 0 0 1-7.749-2.837l-2.544 2.544a1 1 0 0 1-1.414-1.414l2.544-2.544A11.952 11.952 0 0 1 4 16c0-1.476.267-2.89.754-4.197A5 5 0 0 1 8 3ZM5.67 9.89a12.06 12.06 0 0 1 4.22-4.22a3 3 0 0 0-4.22 4.22Zm16.44-4.22a12.06 12.06 0 0 1 4.22 4.22a3 3 0 0 0-4.22-4.22ZM16 10a1 1 0 1 0-2 0v7a1 1 0 0 0 1 1h5a1 1 0 1 0 0-2h-4v-6Z"/>`),
		g.Group(children),
	)
}