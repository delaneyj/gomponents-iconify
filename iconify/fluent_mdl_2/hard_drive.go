package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1856 640q40 0 75 15t61 41t41 61t15 75v576H0V832q0-40 15-75t41-61t61-41t75-15h1664zm64 192q0-26-19-45t-45-19H192q-26 0-45 19t-19 45v448h1792V832zm-256 64h128v128h-128V896zm-256 0h128v128h-128V896z"/>`),
		g.Group(children),
	)
}