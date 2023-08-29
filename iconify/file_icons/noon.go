package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Noon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M127.711 34.453c-170.281 98.17-170.281 344.925 0 443.094S512 452.338 512 256S297.992-63.716 127.711 34.453zM286.662 254.25c0 45.226-62.696 45.226-62.696 0V65.598c0-43.33 62.696-43.33 62.696 0V254.25z"/>`),
		g.Group(children),
	)
}