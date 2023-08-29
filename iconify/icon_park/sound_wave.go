package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundWave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 14V12C6 8.68629 8.68629 6 12 6H36C39.3137 6 42 8.68629 42 12V14"/><path d="M32 18V30"/><path d="M40 20V28"/><path d="M24 15V33"/><path d="M16 18V30"/><path d="M8 20V28"/><path d="M6 34V36C6 39.3137 8.68629 42 12 42H36C39.3137 42 42 39.3137 42 36V34"/></g>`),
		g.Group(children),
	)
}