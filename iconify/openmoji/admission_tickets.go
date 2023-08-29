package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdmissionTickets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#D22F27" d="M57.6 35.9c0-4.9 3.9-9.1 9.4-10.6v-5.2H5v5.5c4.9 1.7 8.4 5.7 8.4 10.3S9.9 44.6 5 46.3v5.5h62v-5.2c-5.4-1.5-9.4-5.7-9.4-10.7z"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-miterlimit="10"><path stroke-width="2.216" d="M57.1 35.6c0-4.9 3.9-9.1 9.4-10.6v-5.2h-62v5.5c4.9 1.7 8.4 5.7 8.4 10.3s-3.5 8.6-8.4 10.3v5.5h62v-5.2c-5.4-1.5-9.4-5.7-9.4-10.6z"/><path stroke-width="2.095" d="M18.5 25h34v20h-34z"/></g><path fill="none" stroke="#FFF" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2.1" d="M18.5 25h34v20h-34z"/>`),
		g.Group(children),
	)
}