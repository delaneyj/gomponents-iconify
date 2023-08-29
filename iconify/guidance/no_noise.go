package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoNoise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M17 6.5a5.5 5.5 0 0 1 .5 10.978M17 9.5a2.5 2.5 0 0 1 0 5m-2.5 0v-12h-.25l-.523.548a15.999 15.999 0 0 1-6.483 4.12l-.06.015M14.5 18v3.5h-.25l-.523-.548A16 16 0 0 0 2.154 16H1.5V8h.654c.79 0 1.574-.058 2.346-.173M.5.5l6.683 6.683M23.5 23.5L7.183 7.183"/>`),
		g.Group(children),
	)
}