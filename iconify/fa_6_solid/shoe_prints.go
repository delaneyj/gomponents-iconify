package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoePrints(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M416 0c-63.7 0-160 32-160 32v128c48 0 76 16 104 32s56 32 104 32c56.4 0 176-16 176-96S512 0 416 0zM128 96c0 35.3 28.7 64 64 64h32V32h-32c-35.3 0-64 28.7-64 64zm160 416c96 0 224-48 224-128s-119.6-96-176-96c-48 0-76 16-104 32s-56 32-104 32v128s96.3 32 160 32zM0 416c0 35.3 28.7 64 64 64h32V352H64c-35.3 0-64 28.7-64 64z"/>`),
		g.Group(children),
	)
}