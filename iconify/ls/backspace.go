package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backspace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M253 94h452c35 0 63 28 63 63v379c0 35-28 64-63 64H253L0 347zm313 408l57-56l-100-99l100-99l-57-57l-99 99l-98-99l-57 56l99 100l-99 98l57 56l99-98z"/>`),
		g.Group(children),
	)
}