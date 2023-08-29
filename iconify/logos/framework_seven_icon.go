package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrameworkSevenIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#EE350F" d="M0 128c0-30.764 10.853-58.992 28.939-81.067H227.06L85.735 248.858C35.811 231.4 0 183.882 0 128Zm237.562-66.219C249.264 81.101 256 103.764 256 128c0 70.692-57.308 128-128 128c-8.47 0-16.746-.823-24.756-2.392L237.562 61.78ZM45.813 29.867C68.05 11.225 96.713 0 128 0c31.286 0 59.951 11.225 82.187 29.867H45.813Z"/>`),
		g.Group(children),
	)
}