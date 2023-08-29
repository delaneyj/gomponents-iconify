package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7 12.119V16l-6-6l6-6v3.966c6.98.164 6.681-4.747 4.904-7.966C16.29 4.741 15.359 12.337 7 12.119z"/>`),
		g.Group(children),
	)
}