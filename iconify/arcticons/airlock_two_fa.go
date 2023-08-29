package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlockTwoFa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.996 15.681c2.896 1.917 4.628 5.096 4.628 8.493h0c0 5.68-4.756 10.287-10.624 10.287c-.557 0-1.114-.043-1.665-.127m-4.397-1.714c-2.858-1.922-4.562-5.078-4.562-8.447c0-5.682 4.756-10.287 10.624-10.287c.573 0 1.145.044 1.71.134m4.164-4.634L32.479 3.5m-17.161 41l2.605-5.886"/>`),
		g.Group(children),
	)
}