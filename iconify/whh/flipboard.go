package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M927.998 320h-288v288q0 13-9.5 22.5t-22.5 9.5h-288v288q0 13-9.5 22.5t-22.5 9.5h-256q-13 0-22.5-9.5t-9.5-22.5V32q0-13 9.5-22.5t22.5-9.5h896q13 0 22.5 9.5t9.5 22.5v256q0 13-9.5 22.5t-22.5 9.5z"/>`),
		g.Group(children),
	)
}