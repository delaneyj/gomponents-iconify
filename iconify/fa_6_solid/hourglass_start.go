package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassStart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32 0C14.3 0 0 14.3 0 32s14.3 32 32 32v11c0 42.4 16.9 83.1 46.9 113.1l67.8 67.9l-67.8 67.9C48.9 353.9 32 394.6 32 437v11c-17.7 0-32 14.3-32 32s14.3 32 32 32h320c17.7 0 32-14.3 32-32s-14.3-32-32-32v-11c0-42.4-16.9-83.1-46.9-113.1L237.3 256l67.9-67.9c30-30 46.9-70.7 46.9-113.1V64c17.7 0 32-14.3 32-32s-14.3-32-32-32H32zm256 437v11H96v-11c0-25.5 10.1-49.9 28.1-67.9l67.9-67.8l67.9 67.9c18 18 28.1 42.4 28.1 67.9z"/>`),
		g.Group(children),
	)
}