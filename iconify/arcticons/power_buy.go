package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerBuy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.317 4.5l.429 6.783l2.627 2.2c.192.161.673-.085.673-.335V6.19l-1.01.05l-2.719-1.74Zm-2.943 11.607c.03.328-.24 1.16-.886.46l-2.933-3.172l-1.235-7.406l3.266 2.416l1.09-.319l.698 8.021Zm5.223 11.51l-.433-4.277l2.038-1.156c2.538-1.44 3.905 1.684.976 3.648l-2.581 1.785Zm.385 4.742l.627 5.953l-8.481 5.188l-5.115-5.283L9.282 15.29l6.854 6.032l13.65-6.978c7.711-3.943 12.667 4.34 5.351 13.104c-1.944 2.33-5.233 3.476-8.155 4.91v.003Z"/>`),
		g.Group(children),
	)
}