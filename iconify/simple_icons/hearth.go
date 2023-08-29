package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hearth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.196 10.318v9.576L.453 24v-9.33zm7.659-4.162l7.692 4.162v9.573L15.853 24v-9.33l-7.658-4.352ZM8.196 0v9.576L.453 13.803V4.155Z"/>`),
		g.Group(children),
	)
}