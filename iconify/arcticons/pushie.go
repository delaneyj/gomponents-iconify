package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pushie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><circle cx="16.78" cy="33.862" r="3.138" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="23.571" cy="26.849" r="3.138" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-miterlimit="7" d="M16.686 30.724v-12.41c0-10.608 18.429-9.117 17.649.865c-.565 7.225-5.217 7.08-7.58 7.418"/>`),
		g.Group(children),
	)
}