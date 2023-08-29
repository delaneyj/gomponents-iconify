package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ivpn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.457 18.573h1.628m0 0l-1.914 10.854m11.038 0l1.913-10.854h3.553a2.971 2.971 0 0 1 2.995 3.645a4.532 4.532 0 0 1-4.28 3.645h-3.553m8.516 3.564l1.914-10.854l5.276 10.854l1.914-10.854m-18.262 0l-5.509 10.854l-1.681-10.854"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1"/>`),
		g.Group(children),
	)
}