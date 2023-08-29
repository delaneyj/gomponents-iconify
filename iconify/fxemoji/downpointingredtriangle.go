package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Downpointingredtriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#FF473E" d="M402.077 165.494L269.825 364.849c-6.562 9.892-21.088 9.892-27.65 0L109.923 165.494c-7.316-11.028.591-25.762 13.825-25.762h264.504c13.234-.001 21.141 14.734 13.825 25.762z"/>`),
		g.Group(children),
	)
}