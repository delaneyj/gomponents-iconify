package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M640 320q0 26-19 45t-45 19H448v1024h128q26 0 45 19t19 45t-19 45l-256 256q-19 19-45 19t-45-19L19 1517q-19-19-19-45t19-45t45-19h128V384H64q-26 0-45-19T0 320t19-45L275 19q19-19 45-19t45 19l256 256q19 19 19 45z"/>`),
		g.Group(children),
	)
}