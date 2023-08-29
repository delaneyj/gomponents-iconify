package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hewlettpackard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m249.27 511.965l52.207-143.487h72.118c12.657 0 26.552-9.731 30.881-21.627l56.848-156.161c9.302-25.557-5.334-46.483-32.539-46.483H328.731L197.277 505.34C-47.775 448.897-79.63 93.484 187.47 9.214L56.71 368.478h54.342l69.426-190.572h40.847l-69.427 190.572l54.322.01l64.698-177.798c9.302-25.557-5.344-46.483-32.52-46.483h-45.683L245.171.035c354.231 2.847 355.934 506.063 4.1 511.93zM371.02 177.77l-57.179 156.883h40.827l57.18-156.883H371.02z"/>`),
		g.Group(children),
	)
}