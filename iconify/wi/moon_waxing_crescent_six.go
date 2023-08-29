package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonWaxingCrescentSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path d="M0 1537q209 0 386-103t280-280t103-386t-103-385.5T386 103T0 0q171 297 171 768q0 239-35.5 430T0 1537z" fill="currentColor"/>`),
		g.Group(children),
	)
}