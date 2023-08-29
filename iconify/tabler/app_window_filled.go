package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppWindowFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M19 4a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3zM6.01 7l-.127.007A1 1 0 0 0 6 9l.127-.007A1 1 0 0 0 6.01 7zm3 0l-.127.007A1 1 0 0 0 9 9l.127-.007A1 1 0 0 0 9.01 7z"/></g>`),
		g.Group(children),
	)
}