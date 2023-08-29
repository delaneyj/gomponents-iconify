package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Package(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m17.078 22.004l-1.758-4.13l-2.007 4.753l-7.52-3.29l.175 3.906l9.437 4.374l10.91-5.365l-.15-4.99l-9.087 4.742zM29.454 6.62L18.52 3.382l-3.005 2.67l-3.09-2.358L1.544 8.2l3.796 3.047l-3.43 5.303l10.88 4.756l2.53-5.998l2.256 5.308l11.393-5.942l-3.105-4.71l3.592-3.345zm-14.177 7.96l-9.06-3.83l9.276-4.102L25.1 9.903l-9.823 4.676z"/>`),
		g.Group(children),
	)
}