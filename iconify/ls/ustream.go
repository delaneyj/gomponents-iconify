package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ustream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M373 534c-71 0-113-55-113-163V29h229v342c0 111-41 163-116 163zm257-172V39c12 9 20 24 20 40v550c0 28-22 50-50 50H50c-28 0-50-22-50-50V79c0-28 22-50 50-50h68v331c0 199 92 288 251 288c163 0 261-92 261-286z"/>`),
		g.Group(children),
	)
}