package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M0 1376V768h2304v608q0 66-47 113t-113 47H160q-66 0-113-47T0 1376zm640-224v128h384v-128H640zm-384 0v128h256v-128H256zM2144 0q66 0 113 47t47 113v224H0V160Q0 94 47 47T160 0h1984z"/>`),
		g.Group(children),
	)
}