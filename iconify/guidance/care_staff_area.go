package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CareStaffArea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M0 22.5h24M13.5 20v-6S11 12.5 7 12.5S.5 14 .5 14v6m17-18.5v3h-3v3h3v3h3v-3h3v-3h-3v-3h-3Zm-10.756 9S3.999 8.752 3.999 6.566c0-1.691 1.344-3.062 3.002-3.062c1.658 0 2.995 1.371 2.995 3.062C9.996 8.75 7.26 10.5 7.26 10.5h-.515Z"/>`),
		g.Group(children),
	)
}