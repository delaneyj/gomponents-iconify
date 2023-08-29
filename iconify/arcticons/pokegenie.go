package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pokegenie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M33.5 5.5h-19a9 9 0 0 0-9 9v19a9 9 0 0 0 9 9h19a9 9 0 0 0 9-9v-19a9 9 0 0 0-9-9Z"/><circle cx="14.484" cy="19.372" r="7.25"/><circle cx="14.484" cy="19.372" r="4"/></g><path fill="none" stroke="currentColor" d="M42.5 22.64H29.794s-2.605-.428-5.599 3.383l-2.024 2.642c-2.993 3.81-5.598 3.382-5.598 3.382H5.5"/>`),
		g.Group(children),
	)
}