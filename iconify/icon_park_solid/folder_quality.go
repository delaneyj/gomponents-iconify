package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderQuality(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFolderQuality0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M5 8a2 2 0 0 1 2-2h12l5 6h17a2 2 0 0 1 2 2v26a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V8Z"/><path fill="#000" stroke="#000" stroke-linecap="round" d="M19.8 21h8.4l2.8 3.53L24 33l-7-8.47L19.8 21Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFolderQuality0)"/>`),
		g.Group(children),
	)
}