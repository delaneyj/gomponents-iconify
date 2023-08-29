package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyOldboy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.188 40.453h0A60.854 60.854 0 0 1 18.181 43.5h0a3.284 3.284 0 0 1-3.294-2.736L9.57 9.327a1.02 1.02 0 0 1 .836-1.175l21.5-3.638a1.02 1.02 0 0 1 1.175.836L38.4 36.786a3.284 3.284 0 0 1-2.21 3.667Zm-16.463-5.878l-.851-5.027m-2.088 2.939l5.027-.851m1.208 7.143l2.03-.343m2.998-.508l2.03-.343"/><circle cx="29.751" cy="31.063" r=".75" fill="currentColor"/><circle cx="33.003" cy="28.974" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.247 25.155l1.512-.083A92.16 92.16 0 0 0 34.3 21.58l1.29-.374M13.39 10.83l16.577-2.804l1.879 11.107l-16.577 2.804z"/>`),
		g.Group(children),
	)
}