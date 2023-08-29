package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaymentMethod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M38 10L10 38"/><path d="M6 6L12 14L18 6"/><path d="M5 14H19"/><path d="M5 20H19"/><path d="M12 14V26"/><path d="M32.8462 26H42V42H22V36.15"/></g>`),
		g.Group(children),
	)
}