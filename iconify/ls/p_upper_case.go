package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 0h279c58 0 109 27 141 71c22 30 35 66 35 105s-13 75-35 104c-32 44-83 72-141 72H72v383H0V0zm72 280h210c56-1 101-47 101-104c0-56-45-103-101-105H72v209z"/>`),
		g.Group(children),
	)
}