package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewspaperFolding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#000" d="M22 44L21 36"/><path fill="#2F88FF" stroke="#000" d="M42 44V12H26L27 20L28 28L29 36L22 44H42Z"/><path stroke="#fff" d="M28 28H33"/><path stroke="#fff" d="M27 20H33"/><path fill="#2F88FF" stroke="#000" d="M6 4H25L26 12L27 20L28 28L29 36H21H6V4Z"/><path stroke="#fff" d="M12 12H19"/><path stroke="#fff" d="M12 20H20"/><path stroke="#fff" d="M12 28H21"/></g>`),
		g.Group(children),
	)
}