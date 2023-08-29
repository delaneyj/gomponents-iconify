package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yaml(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M342.016 0H457L114.983 512H0l171.008-256L0 0h114.983L228.5 169.934z"/>`),
		g.Group(children),
	)
}