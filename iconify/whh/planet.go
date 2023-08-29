package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Planet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M830 359q-14 119-104.5 200T513 640q-79 0-149-37q-128 40-224.5 37T13 588q-32-50 19-133.5T196 282q14-120 104.5-200.5T513 1q79 0 149 37Q790-2 886.5 1T1013 52q32 51-19 134T830 359zM137 542q15 27 61 33t110-10q-87-72-108-183q-94 106-63 160zM889 98q-15-27-61-32.5T718 75q87 73 108 183q94-105 63-160z"/>`),
		g.Group(children),
	)
}