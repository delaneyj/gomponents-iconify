package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Syncopoli(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.52 20.8c8.69-7.51 16.29-11.9 23.55-11.9c6.58 0 8.79 5.18 7.4 9.18"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.35 15.55l12.15 5.02l-11.28 5.29l-.87-10.31zM41.48 27.2c-8.69 7.51-16.29 11.9-23.55 11.9c-6.58 0-8.79-5.18-7.4-9.18"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.65 32.45L4.5 27.43l11.28-5.29l.87 10.31z"/>`),
		g.Group(children),
	)
}