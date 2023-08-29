package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowerBud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M13 10a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/><path fill-rule="evenodd" d="M15.944 13.371a3.502 3.502 0 0 0 0-6.742A3.5 3.5 0 0 0 10 3.55a3.5 3.5 0 0 0-5.944 3.078a3.502 3.502 0 0 0 0 6.742A3.5 3.5 0 0 0 10 16.45a3.5 3.5 0 0 0 5.944-3.078Zm-1.091-6.523a.5.5 0 0 0 .417.666a2.5 2.5 0 0 1 0 4.972a.5.5 0 0 0-.417.666a2.5 2.5 0 0 1-4.436 2.23a.5.5 0 0 0-.833 0a2.5 2.5 0 0 1-4.436-2.23a.5.5 0 0 0-.418-.666a2.5 2.5 0 0 1 0-4.972a.5.5 0 0 0 .417-.666a2.5 2.5 0 0 1 4.436-2.23a.5.5 0 0 0 .833 0a2.5 2.5 0 0 1 4.436 2.23Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}