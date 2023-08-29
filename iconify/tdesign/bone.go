package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 3a1.5 1.5 0 0 0-1.318 2.217l.358.657L5.874 16.54l-.657-.358a1.5 1.5 0 1 0-.993 2.793l.676.125l.125.676a1.5 1.5 0 1 0 2.793-.993l-.358-.657L18.126 7.46l.657.358a1.5 1.5 0 1 0 .993-2.793L19.1 4.9l-.125-.676A1.5 1.5 0 0 0 17.5 3ZM14 4.5a3.5 3.5 0 0 1 6.764-1.264a3.5 3.5 0 0 1-2.218 6.632l-8.678 8.678a3.5 3.5 0 0 1-6.633 2.218a3.5 3.5 0 0 1 2.219-6.632l8.678-8.678A3.502 3.502 0 0 1 14 4.5Z"/>`),
		g.Group(children),
	)
}