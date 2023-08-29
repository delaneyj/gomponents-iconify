package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Instalod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M153.384 480h233.729l115.441-204.235l-298.325 57.446Zm351.342-239.922L387.113 32H155.669L360.23 267.9ZM124.386 48.809L7.274 256l115.962 205.154l102.391-295.593Z"/>`),
		g.Group(children),
	)
}