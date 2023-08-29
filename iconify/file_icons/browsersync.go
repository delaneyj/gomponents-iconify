package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Browsersync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M159.27.793L4.263 78.296C1.756 79.55 0 82.058 0 85.068v341.864c0 2.76 1.505 5.518 4.264 6.772l155.005 77.503c5.016 2.508 10.785-1.254 10.785-6.772V7.565c0-5.518-5.769-9.28-10.785-6.772zm41.46 510.414l155.006-77.503c2.508-1.254 4.264-3.762 4.264-6.772V260.64c0-2.759-1.505-5.518-4.264-6.772l-155.005-77.503c-5.016-2.508-10.785 1.255-10.785 6.773v321.297c0 5.518 5.769 9.28 10.785 6.772z"/>`),
		g.Group(children),
	)
}