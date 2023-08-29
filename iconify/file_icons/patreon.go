package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Patreon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 501.505h90.01V10.5H0v491.005zm235.696-147.654c-122.432-70.584-122.432-248 0-318.585S512 53.391 512 194.56S358.129 424.434 235.696 353.85z"/>`),
		g.Group(children),
	)
}