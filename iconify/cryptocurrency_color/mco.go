package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mco(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#103F68"/><path fill="#FFF" fill-rule="nonzero" d="m15.98 5.018l9.52 5.483v11L15.991 27l-.077-.019l-9.414-5.48v-11l9.414-5.483h.066zm-.031 1.138L7.5 11.076v9.85l8.448 4.919l1.032-.597l7.52-4.325v-9.845l-7.52-4.35l-1.031-.572zm-7.14 10.61l2.501-1.87l2.211 1.412v2.54l1.673 1.612l-.001.756l-1.612 1.51H12.22l-3.41-5.96zm7.903 4.452l-.003-.76l1.667-1.61v-2.54l2.187-1.43l2.496 1.889l-3.393 5.942h-1.344l-1.61-1.491zm-2.37-4.91l-.814-2.131h4.838l-.798 2.131l.236 2.382l-1.867.004l-1.845.003l.25-2.389zm1.595-2.715l-4.598-.002l.855-3.82h7.464l.9 3.825l-4.621-.003z"/></g>`),
		g.Group(children),
	)
}