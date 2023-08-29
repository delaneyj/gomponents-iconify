package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lightningmood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#FDB436" d="m479.23 258.78l-197.7-36.571c-4.891-.905-7.09-6.667-4.044-10.6L427.065 18.432c4.498-5.808-2.366-13.542-8.667-9.766L21.336 246.595c-5.216 3.126-3.84 11.036 2.126 12.216l180.068 35.606c4.694.928 6.897 6.365 4.173 10.299L77.05 493.386c-4.05 5.849 2.684 13.111 8.821 9.513l395.511-231.856c5.308-3.111 3.899-11.144-2.152-12.263z"/>`),
		g.Group(children),
	)
}