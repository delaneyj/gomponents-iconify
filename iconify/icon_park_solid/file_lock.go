package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFileLock0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M10 44h28a2 2 0 0 0 2-2V14H30V4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2Z"/><path stroke="#fff" d="m30 4l10 10"/><path fill="#000" stroke="#000" d="M17 27h14v8H17z"/><path stroke="#000" d="M28 27v-4c0-1.657-1-4-4-4s-4 2.343-4 4v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFileLock0)"/>`),
		g.Group(children),
	)
}