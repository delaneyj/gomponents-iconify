package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CookingPot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M10 44H38V20.9474C38 14.9013 31.732 10 24 10C16.268 10 10 14.9013 10 20.9474V44Z" clip-rule="evenodd"/><path fill="#2F88FF" d="M38 22.0437C38 21.8001 38 21.4346 38 20.9474C38 14.9013 31.732 10 24 10C16.268 10 10 14.9013 10 20.9474C10 21.4332 10 21.7975 10 22.0404L38 22.0437Z"/><path stroke-linecap="round" d="M4 22H44"/><path stroke-linecap="round" d="M21 4H27"/></g>`),
		g.Group(children),
	)
}