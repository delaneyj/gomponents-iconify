package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicRhythm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M5 42H10"/><path d="M5 36H10"/><path d="M5 30H10"/><path d="M5 24H10"/><path d="M16 42H21"/><path d="M16 36H21"/><path d="M16 30H21"/><path d="M16 24H21"/><path d="M16 18H21"/><path d="M16 12H21"/><path d="M16 6H21"/><path d="M27 42H32"/><path d="M38 42H43"/><path d="M27 36H32"/><path d="M38 36H43"/><path d="M27 30H32"/><path d="M38 30H43"/><path d="M38 24H43"/><path d="M38 18H43"/></g>`),
		g.Group(children),
	)
}