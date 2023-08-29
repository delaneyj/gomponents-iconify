package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrebelMusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.546 32.84V5.5h-8.078v29.544h0a7.457 7.457 0 0 0 7.457 7.456h5.705v-7.457h-2.88a2.203 2.203 0 0 1-2.204-2.202Zm11.597-14.679l-8.05-4.647v9.295l8.05-4.648zM14.91 13.514v9.295m-5.053-9.295v9.295"/>`),
		g.Group(children),
	)
}