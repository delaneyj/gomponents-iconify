package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeIndesign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M265.301 269.334c-1.24-39.937 26.821-71.547 67.562-60.916v122.94c-40.599 9.808-69.274-14.325-67.562-62.024zM0 0v512h512V0H0zm177.804 369.014H134.61v-266.92h43.194v266.92zm135.122 3.322c-49.296 1.424-92.89-29.841-91.927-100.787c.774-57.015 36.303-108.212 111.864-103.002v-66.453h44.302v239.785c0 4.067.37 9.887 1.363 14.713c-17.01 9.418-39.047 15.574-65.602 15.744z"/>`),
		g.Group(children),
	)
}