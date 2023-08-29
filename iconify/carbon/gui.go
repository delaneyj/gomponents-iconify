package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gui(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="20" cy="8" r="1" fill="currentColor"/><circle cx="23" cy="8" r="1" fill="currentColor"/><circle cx="26" cy="8" r="1" fill="currentColor"/><path fill="currentColor" d="M28 4H4a2.002 2.002 0 0 0-2 2v20a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2V6a2.002 2.002 0 0 0-2-2Zm0 2v4H4V6ZM4 12h6v14H4Zm8 14V12h16v14Z"/>`),
		g.Group(children),
	)
}