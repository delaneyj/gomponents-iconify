package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EngineBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M244 112a12 12 0 0 0-12 12v16h-16v-8a20 20 0 0 0-20-20h-11l-34.17-34.14A19.86 19.86 0 0 0 136.69 72H124V56h16a12 12 0 0 0 0-24H84a12 12 0 0 0 0 24h16v16H60a20 20 0 0 0-20 20v48H24v-16a12 12 0 0 0-24 0v56a12 12 0 0 0 24 0v-16h16v12.69a19.86 19.86 0 0 0 5.86 14.14l35.31 35.31A19.86 19.86 0 0 0 95.31 232h41.38a19.86 19.86 0 0 0 14.14-5.86L185 192h11a20 20 0 0 0 20-20v-8h16v16a12 12 0 0 0 24 0v-56a12 12 0 0 0-12-12Zm-52 56h-8.69a19.86 19.86 0 0 0-14.14 5.86L135 208H97l-33-33V96h71l34.14 34.14a19.86 19.86 0 0 0 14.17 5.86H192Z"/>`),
		g.Group(children),
	)
}