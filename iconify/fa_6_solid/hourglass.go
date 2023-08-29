package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 32C0 14.3 14.3 0 32 0h320c17.7 0 32 14.3 32 32s-14.3 32-32 32v11c0 42.4-16.9 83.1-46.9 113.1L237.3 256l67.9 67.9c30 30 46.9 70.7 46.9 113.1v11c17.7 0 32 14.3 32 32s-14.3 32-32 32H32c-17.7 0-32-14.3-32-32s14.3-32 32-32v-11c0-42.4 16.9-83.1 46.9-113.1l67.8-67.9l-67.8-67.9C48.9 158.1 32 117.4 32 75V64C14.3 64 0 49.7 0 32zm96 32v11c0 25.5 10.1 49.9 28.1 67.9l67.9 67.8l67.9-67.9c18-18 28.1-42.4 28.1-67.9V64H96zm0 384h192v-11c0-25.5-10.1-49.9-28.1-67.9L192 301.3l-67.9 67.9c-18 18-28.1 42.4-28.1 67.9v11z"/>`),
		g.Group(children),
	)
}