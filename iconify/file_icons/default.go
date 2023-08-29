package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Default(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M288 32H0v448h384V128l-96-96zm64 416H32V64h224l96 96v288z"/>`),
		g.Group(children),
	)
}