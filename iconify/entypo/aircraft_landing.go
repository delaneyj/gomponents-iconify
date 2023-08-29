package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AircraftLanding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.752 16.038c-.097.266-.822 1.002-6.029-.878l-5.105-1.843C5.841 12.676 3.34 11.668 2.36 11.1c-.686-.397-.836-1.282-.836-1.282s-.163-2.956-.263-3.684c-.1-.728.095-.853.796-.492c.436.225 1.865 2.562 2.464 3.567c1.512.381 2.862.761 3.493.949c-.257-1.717-.74-4.928-.913-5.933c-.166-.963.55-.535.55-.535c.331.19.983.661 1.206 1.002c1.522 2.326 3.672 6.6 3.836 6.928c.896.28 2.277.733 3.102 1.03c2.156.779 3.087 3.034 2.957 3.388z"/>`),
		g.Group(children),
	)
}