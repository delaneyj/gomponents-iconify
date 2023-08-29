package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileMusicOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFileMusicOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 23v-9L31 4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h12m18-15l-4 .969V40.5"/><circle cx="32" cy="40" r="4" fill="#fff"/><path d="M30 4v10h10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFileMusicOne0)"/>`),
		g.Group(children),
	)
}