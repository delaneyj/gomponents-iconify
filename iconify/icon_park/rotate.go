package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rotate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M12 24H42V42H12V24Z"/><path stroke-linecap="round" d="M6 8V17H15"/><path stroke-linecap="round" d="M38.4747 13.2985C35.1956 8.87049 29.933 6 24 6C18.1788 6 13.0029 8.76334 9.71272 13.0498L6 17"/></g>`),
		g.Group(children),
	)
}