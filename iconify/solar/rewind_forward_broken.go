package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindForwardBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m11 15.232l-6.097 4.46C3.601 20.589 2 19.422 2 17.574V15m9-6.232l-6.097-4.46C3.601 3.411 2 4.578 2 6.426V11m15.37-3.292L13.66 5.27C12.467 4.485 11 5.507 11 7.123v9.754c0 1.616 1.467 2.638 2.661 1.853l7.418-4.877c1.228-.807 1.228-2.899 0-3.706l-.928-.61"/>`),
		g.Group(children),
	)
}