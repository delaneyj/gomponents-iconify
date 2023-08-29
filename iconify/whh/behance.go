package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Behance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M640 64h320v64H640V64zm160.5 128q104.5 0 164 73t59.5 183H640q0 85 40.5 138.5T800 640q63 0 101-35t52-93h65q-16 86-72.5 139T800 704q-104 0-164-73t-60-183t60-183t164.5-73zM953 384q-13-59-51.5-93.5t-101-34.5t-101 34.5T647 384h306zM512 512q0 79-62.5 135.5T298 704H0V0h261q77 0 132 56t55 136q0 81-51 140q49 21 82 70.5T512 512zM264 64H64v256h200q50 0 85-37.5t35-90.5t-35-90.5T264 64zm40 320H64v256h240q59 0 101.5-37.5T448 512t-42.5-90.5T304 384z"/>`),
		g.Group(children),
	)
}