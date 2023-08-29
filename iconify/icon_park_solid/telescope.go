package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Telescope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTelescope0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="12" height="38" x="6" y="5" rx="6"/><rect width="12" height="38" x="30" y="5" rx="6"/><path fill="#fff" d="M12 43a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm24 0a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path stroke-linecap="round" d="M30 21a6 6 0 0 0-12 0m12 10a6 6 0 0 0-12 0"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTelescope0)"/>`),
		g.Group(children),
	)
}