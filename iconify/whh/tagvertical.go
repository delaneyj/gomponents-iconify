package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tagvertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M641 1024H129q-53 0-90.5-37.5T1 896V393q-2-29 18-50L340 19q19-19 45-19q27 0 46 19l80 81l160 162l80 81q9 10 14 20.5t4 16.5v516q0 53-37.5 90.5T641 1024zM385 192q-53 0-90.5 37.5T257 320t37.5 90.5T385 448t90.5-37.5T513 320t-37.5-90.5T385 192z"/>`),
		g.Group(children),
	)
}