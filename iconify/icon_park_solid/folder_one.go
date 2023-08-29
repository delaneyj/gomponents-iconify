package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFolderOne0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M5 6a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v36a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6Z"/><circle cx="11" cy="35" r="2" fill="#000"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M11 11v14"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m23.977 9.756l10.06-2.515a1 1 0 0 1 1.211.72l7.5 29.063a1 1 0 0 1-.725 1.22l-10.06 2.515a1 1 0 0 1-1.211-.72l-7.5-29.063a1 1 0 0 1 .726-1.22Z"/><circle cx="35" cy="32" r="2" fill="#000"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m31 16l2.5 10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFolderOne0)"/>`),
		g.Group(children),
	)
}