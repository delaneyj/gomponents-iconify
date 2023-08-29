package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmDbTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="23" cy="26" r="1" fill="currentColor"/><circle cx="9" cy="6" r="1" fill="currentColor"/><path fill="currentColor" d="M12 17h8c2.206 0 4-1.794 4-4v-3h2c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2H6c-1.103 0-2 .897-2 2v4c0 1.103.897 2 2 2h16v3c0 1.103-.897 2-2 2h-8c-2.206 0-4 1.794-4 4v3H6c-1.103 0-2 .897-2 2v4c0 1.103.897 2 2 2h20c1.103 0 2-.897 2-2v-4c0-1.103-.897-2-2-2H10v-3c0-1.103.897-2 2-2ZM6 4h20v4H6V4Zm20 24H6v-4h20v4Z"/>`),
		g.Group(children),
	)
}