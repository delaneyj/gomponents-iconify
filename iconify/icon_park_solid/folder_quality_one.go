package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderQualityOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFolderQualityOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M43 23v-9a2 2 0 0 0-2-2H24l-5-6H7a2 2 0 0 0-2 2v32a2 2 0 0 0 2 2h15"/><path fill="#fff" d="M29.8 29h8.4l2.8 3.53L34 41l-7-8.47L29.8 29Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFolderQualityOne0)"/>`),
		g.Group(children),
	)
}