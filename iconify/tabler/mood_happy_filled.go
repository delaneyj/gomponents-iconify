package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodHappyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M17 3.34a10 10 0 1 1-14.995 8.984L2 12l.005-.324A10 10 0 0 1 17 3.34zM15 13H9a1 1 0 0 0-1 1v.05a3.975 3.975 0 0 0 3.777 3.97l.227.005a4.026 4.026 0 0 0 3.99-3.79l.006-.206A1 1 0 0 0 15 13zM9.01 8l-.127.007A1 1 0 0 0 9 10l.127-.007A1 1 0 0 0 9.01 8zm6 0l-.127.007A1 1 0 0 0 15 10l.127-.007A1 1 0 0 0 15.01 8z"/></g>`),
		g.Group(children),
	)
}