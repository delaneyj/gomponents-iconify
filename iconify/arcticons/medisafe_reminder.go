package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedisafeReminder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.768 15.836H11.944c-4.349 0-8.146 3.281-8.426 7.62a8.165 8.165 0 0 0 8.146 8.708h2.274a6.044 6.044 0 0 0 5.73-4.12l4.1-12.208Zm12.291 0h-3.13c-2.122 0-3.561 1.039-4.08 2.581l-4.616 13.747h12.103a8.165 8.165 0 0 0 8.146-8.707c-.28-4.34-4.075-7.621-8.423-7.621Z"/>`),
		g.Group(children),
	)
}