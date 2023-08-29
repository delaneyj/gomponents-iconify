package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TruckLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m245.57 117.78l-14-35a13.93 13.93 0 0 0-13-8.8H182V64a6 6 0 0 0-6-6H24a14 14 0 0 0-14 14v112a14 14 0 0 0 14 14h18.6a30 30 0 0 0 58.8 0h53.2a30 30 0 0 0 58.8 0H232a14 14 0 0 0 14-14v-64a6 6 0 0 0-.43-2.22ZM182 86h36.58a2 2 0 0 1 1.86 1.26l10.7 26.74H182ZM22 72a2 2 0 0 1 2-2h146v68H22Zm50 138a18 18 0 1 1 18-18a18 18 0 0 1-18 18Zm82.6-24h-53.2a30 30 0 0 0-58.8 0H24a2 2 0 0 1-2-2v-34h148v15.48A30.1 30.1 0 0 0 154.6 186Zm29.4 24a18 18 0 1 1 18-18a18 18 0 0 1-18 18Zm50-26a2 2 0 0 1-2 2h-18.6a30.05 30.05 0 0 0-29.4-24c-.67 0-1.34 0-2 .07V126h52Z"/>`),
		g.Group(children),
	)
}