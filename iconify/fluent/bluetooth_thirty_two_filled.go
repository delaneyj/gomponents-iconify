package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.25 2.104a1.25 1.25 0 0 1 1.35.23l7 6.5a1.25 1.25 0 0 1-.018 1.848L17.627 16l5.956 5.318a1.25 1.25 0 0 1 .018 1.848l-7 6.5a1.25 1.25 0 0 1-2.101-.916v-9.958l-4.917 4.39a1.25 1.25 0 1 1-1.666-1.864L13.873 16l-5.956-5.318a1.25 1.25 0 1 1 1.666-1.864l4.917 4.39V3.25c0-.497.295-.947.75-1.146ZM17 18.792v7.091l3.893-3.615L17 18.792Zm0-5.584l3.893-3.476L17 6.117v7.091Z"/>`),
		g.Group(children),
	)
}