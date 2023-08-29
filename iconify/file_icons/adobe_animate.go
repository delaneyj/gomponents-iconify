package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeAnimate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m169.658 137.076l27.928 109.382h-55.157l27.23-109.382zM512 0v512H0V0h512zM271.325 356.538l-75.825-256h-52.52c-.531 13.05-1.326 15.964-2.645 20.713L73.008 356.538h40.468l19.179-71.68h74.705l19.766 71.68h44.2zm162.246-124.043c0-73.454-63.051-83.034-142.537-52.739c1.164 13.499 1.504 23.648 1.504 25.975v151.04h43.055V205.265c34.821-9.685 57.917-4.355 57.25 30.488v120.785h40.728V232.495z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}