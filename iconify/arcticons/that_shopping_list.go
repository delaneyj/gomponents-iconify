package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThatShoppingList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.61 19.999h-7.527l-7.43-12.386c-.382-.571-.954-.953-1.716-.953s-1.334.382-1.715.953L14.792 20H7.454c-2 0-3.334 1.905-2.858 3.715L8.789 39.34c.286 1.143 1.429 2 2.572 2h25.247a2.659 2.659 0 0 0 2.573-2l4.192-15.53c.572-1.905-.858-3.81-2.763-3.81Zm-16.578-7.622l4.573 7.622h-9.05l4.477-7.622Zm-9.05 7.622h18.101"/><circle cx="13.552" cy="26.096" r="3.906" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24.032" cy="26.096" r="3.906" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.512" cy="26.096" r="3.906" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.035" cy="35.623" r="3.906" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="18.316" cy="35.623" r="3.906" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.557 25.143l-2.001 2.001c-.095.095-.19.095-.286 0l-.953-.953"/>`),
		g.Group(children),
	)
}