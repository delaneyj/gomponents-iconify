package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BalloonFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 1a7 7 0 0 1 7 7c0 5.457-3.028 10-7 10c-3.9 0-6.89-4.379-6.997-9.703L5 8l.004-.24A7 7 0 0 1 12 1zm0 4a1 1 0 0 0 0 2l.117.007A1 1 0 0 1 13 8l.007.117A1 1 0 0 0 15 8a3 3 0 0 0-3-3z"/><path fill="currentColor" d="M12 16a1 1 0 0 1 .993.883L13 17v1a3 3 0 0 1-2.824 2.995L10 21H7a1 1 0 0 0-.993.883L6 22a1 1 0 0 1-2 0a3 3 0 0 1 2.824-2.995L7 19h3a1 1 0 0 0 .993-.883L11 18v-1a1 1 0 0 1 1-1z"/></g>`),
		g.Group(children),
	)
}