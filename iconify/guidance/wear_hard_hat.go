package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WearHardHat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3.52 13c.105 2.458.66 4.304 2.438 6.244C7.658 21.1 9.89 22.688 12 23.25c2.112-.563 4.35-2.155 6.052-4.015c1.77-1.935 2.324-3.783 2.428-6.235M8.5 6l.799-4.17M15.5 6l-.799-4.17m0 0L14.5.78A10.903 10.903 0 0 0 12 .5c-.864 0-1.704.093-2.5.28l-.201 1.05m5.402 0c2.804.725 5.033 2.65 5.799 5.867c.261 1.096.563 2.178 1.5 2.803v.5s-2.257 1-10 1s-10-1-10-1v-.5c.937-.625 1.239-1.707 1.5-2.803c.766-3.216 2.995-5.142 5.799-5.867"/>`),
		g.Group(children),
	)
}