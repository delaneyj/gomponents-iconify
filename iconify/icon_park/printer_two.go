package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" d="M38 20V8C38 6.89543 37.1046 6 36 6H12C10.8954 6 10 6.89543 10 8V20"/><rect width="36" height="22" x="6" y="20" rx="2"/><path fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round" d="M20 34H35V42H20V34Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 26H15"/></g>`),
		g.Group(children),
	)
}