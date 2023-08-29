package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blacklistblocker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.334 34.041v7.71a1.5 1.5 0 0 1-1.5 1.5H11.52a3.854 3.854 0 0 1-3.855-3.855h0a3.854 3.854 0 0 1 3.855-3.854h27.313a1.5 1.5 0 0 0 1.5-1.5V6.25a1.5 1.5 0 0 0-1.5-1.5H11.52a3.854 3.854 0 0 0-3.855 3.854v30.792"/><ellipse cx="24.125" cy="19.953" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="11.271" ry="9.336"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.554 20.146c.503-.426 2.392-2.028 2.568-2.169a.734.734 0 0 0 .188-.811a13.86 13.86 0 0 1-.36-1.696a.638.638 0 0 0-.724-.478l-.023.002h-2.757a.449.449 0 0 0-.44.374a8.313 8.313 0 0 0 1.495 4.7v.001l.012.017l.04.06l.003-.002a10.745 10.745 0 0 0 4.22 3.489l-.004.003l.141.065l.005.002h0a13.337 13.337 0 0 0 5.617 1.215a.454.454 0 0 0 .451-.364v-2.286a.569.569 0 0 0-.553-.617l-.02-.002q-1.032-.105-2.046-.298a1.181 1.181 0 0 0-.979.156c-.17.146-2.1 1.71-2.616 2.129m-7.615 2.918l15.936-13.202"/>`),
		g.Group(children),
	)
}