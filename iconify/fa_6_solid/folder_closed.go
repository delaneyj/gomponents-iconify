package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 480H64c-35.3 0-64-28.7-64-64V192h512v224c0 35.3-28.7 64-64 64zm64-320H0V96c0-35.3 28.7-64 64-64h128c20.1 0 39.1 9.5 51.2 25.6l19.2 25.6c6 8.1 15.5 12.8 25.6 12.8h160c35.3 0 64 28.7 64 64z"/>`),
		g.Group(children),
	)
}