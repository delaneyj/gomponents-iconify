package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPaint0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m15.536 22.898l9.899 9.9m-9.899-9.9L7.05 31.383a7 7 0 1 0 9.9 9.9l8.485-8.486m-9.899 0l-4.243 4.243"/><path fill="#555" d="m25.435 32.797l14.907-6.432c2.688-1.16 3.809-4.379 2.086-6.745C38.264 13.903 32.65 8.89 28.51 5.823c-2.29-1.696-5.33-.64-6.46 1.975l-6.514 15.1l9.899 9.9Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPaint0)"/>`),
		g.Group(children),
	)
}