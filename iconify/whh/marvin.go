package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Marvin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 449q-114 0-266.5 14.5T0 495q3-101 45-192.5t110-158T316.5 39T512 0t195.5 39T869 144.5t110 158t45 192.5q-93-17-245.5-31.5T512 449zM223 641l150-124q77-4 139-4t139 4l150 124l64-108q98 11 158 23q-11 130-81.5 237.5t-184 170T512 1026t-246-62.5t-184-170T1 556q60-12 158-23z"/>`),
		g.Group(children),
	)
}