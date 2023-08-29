package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteThemes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDeleteThemes0"><g fill="none" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke="#fff" stroke-linejoin="round" d="M8 15h32l-3 29H11L8 15Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" d="M20.002 25.002v10m8-10.002v9.997"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M12 15L28.324 3L36 15"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDeleteThemes0)"/>`),
		g.Group(children),
	)
}