package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1088 832V192H640v1137q119-63 213-137q235-184 235-360zm192-768v768q0 86-33.5 170.5t-83 150t-118 127.5T919 1383t-121 77.5t-89.5 49.5t-42.5 20q-12 6-26 6t-26-6q-16-7-42.5-20t-89.5-49.5t-121-77.5t-126.5-103t-118-127.5t-83-150T0 832V64q0-26 19-45T64 0h1152q26 0 45 19t19 45z"/>`),
		g.Group(children),
	)
}