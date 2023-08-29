package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoHearts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#FF473E" d="M46.3 34.6c-5.8-5.9-15.2-5.5-20.5.7c-.7.8-2 .9-2.8.2c-5.8-5-14.6-4.4-19.6 1.4c-4.3 4.9-4.5 12.4-.5 17.6c1.2 1.6 2.7 2.8 4.4 3.7l18.9 13.1c.7.5 1.6.4 2.2-.2L45 55.3c1.4-1.1 2.7-2.5 3.6-4.2c2.9-5.3 2-12.1-2.3-16.5z"/><path fill="#FF7DDA" d="M70.4 7c-3.3-5.2-10.2-6.3-15-2.7c-.7.5-1.6.4-2.1-.3c-3.5-4.5-9.9-5.4-14.5-2c-3.8 3-5.1 8.4-3 12.8c.7 1.4 1.5 2.5 2.6 3.4l11.7 12.4c.4.4 1.1.5 1.6.2l14.5-8.9c1.2-.6 2.3-1.4 3.3-2.5c3-3.5 3.3-8.6.9-12.4z"/>`),
		g.Group(children),
	)
}