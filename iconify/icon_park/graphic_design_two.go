package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphicDesignTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M15.6496 28.2381C12.221 25.6863 10 21.6025 10 17C10 9.26801 16.268 3 24 3C31.732 3 38 9.26801 38 17C38 21.5352 35.8435 25.5668 32.5 28.1252"/><path fill="#2F88FF" fill-rule="evenodd" d="M24 17L44 44H4L24 17Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}