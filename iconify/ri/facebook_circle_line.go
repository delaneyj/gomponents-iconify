package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacebookCircleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.001 19.938a8.001 8.001 0 0 0-1-15.938a8 8 0 0 0-1 15.938V14h-2v-2h2v-1.654c0-1.337.14-1.822.4-2.311A2.725 2.725 0 0 1 12.537 6.9c.382-.205.857-.328 1.687-.381c.329-.021.755.005 1.278.08v1.9h-.5c-.917 0-1.296.043-1.522.164a.728.728 0 0 0-.314.314c-.12.226-.164.45-.164 1.368V12h2.5l-.5 2h-2v5.938Zm-1 2.062c-5.523 0-10-4.477-10-10s4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10Z"/>`),
		g.Group(children),
	)
}