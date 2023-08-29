package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindForwardBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 17.574V6.426C2 4.58 3.6 3.411 4.903 4.308L11 8.768V7.123c0-1.616 1.467-2.638 2.661-1.853l7.417 4.877c1.229.807 1.229 2.899 0 3.706l-7.417 4.877c-1.194.785-2.661-.237-2.661-1.853v-1.645l-6.097 4.46C3.601 20.589 2 19.422 2 17.574Z"/>`),
		g.Group(children),
	)
}