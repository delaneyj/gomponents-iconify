package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Voltage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M315 832h133l-192 192L64 832h123l142-256H0L192 0h128L171 448h341z"/>`),
		g.Group(children),
	)
}