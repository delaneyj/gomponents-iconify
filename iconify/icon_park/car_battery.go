package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarBattery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M43 12H3V42H43V12Z"/><path fill="#2F88FF" stroke="#000" d="M11 6H3V12H11V6Z"/><path fill="#2F88FF" stroke="#000" d="M43 6H35V12H43V6Z"/><path stroke="#fff" d="M9 21H15"/><path stroke="#fff" d="M31 21H37"/><path stroke="#fff" d="M12 18V24"/></g>`),
		g.Group(children),
	)
}