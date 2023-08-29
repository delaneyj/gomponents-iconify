package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6.344 12l2.121 2.122m.707-4.95l2.121 2.12M12 6.344l2.122 2.122M5.07 13.273l8.202-8.203c.792-.792 1.188-1.188 1.645-1.336a2 2 0 0 1 1.236 0c.456.148.852.544 1.643 1.335L18.93 6.2c.792.792 1.188 1.19 1.337 1.646c.13.402.13.834 0 1.235c-.149.457-.545.853-1.337 1.645l-8.202 8.203c-.792.792-1.189 1.188-1.645 1.336c-.402.13-.834.13-1.235 0c-.457-.148-.854-.544-1.646-1.336l-1.133-1.133c-.79-.791-1.187-1.187-1.335-1.643a2 2 0 0 1 0-1.236c.148-.457.545-.853 1.337-1.645Z"/>`),
		g.Group(children),
	)
}