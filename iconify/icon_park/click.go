package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Click(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 4V12"/><path fill="#2F88FF" fill-rule="evenodd" d="M22 22L42 26L36 30L42 36L36 42L30 36L26 42L22 22Z" clip-rule="evenodd"/><path d="M38.1421 9.85789L32.4853 15.5147"/><path d="M9.85787 38.1421L15.5147 32.4853"/><path d="M4 24H12"/><path d="M9.85795 9.85787L15.5148 15.5147"/></g>`),
		g.Group(children),
	)
}