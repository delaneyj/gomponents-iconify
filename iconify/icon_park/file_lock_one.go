package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileLockOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 23V14L31 4H10C8.89543 4 8 4.89543 8 6V42C8 43.1046 8.89543 44 10 44H22"/><rect width="14" height="8" x="28" y="34" fill="#2F88FF"/><path d="M38 34V31C38 29.3431 36.6569 28 35 28C33.3431 28 32 29.3431 32 31V34"/><path d="M30 4V14H40"/></g>`),
		g.Group(children),
	)
}