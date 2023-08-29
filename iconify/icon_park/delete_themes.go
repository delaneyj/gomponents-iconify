package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteThemes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" fill-rule="evenodd" stroke="#000" stroke-linejoin="round" d="M8 15H40L37 44H11L8 15Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" d="M20.002 25.0024V35.0026"/><path stroke="#fff" stroke-linecap="round" d="M28.0024 24.9995V34.9972"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M12 14.9999L28.3242 3L36 15"/></g>`),
		g.Group(children),
	)
}