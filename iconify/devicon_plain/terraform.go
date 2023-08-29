package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terraform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M46.324 26.082L77.941 44.5v36.836L46.324 62.918zM81.41 44.5v36.836l31.633-18.418V26.082zM11.242 5.523V42.36L42.86 60.777V23.941zm66.699 79.852L46.324 66.957v36.824L77.941 122.2zm0 0"/>`),
		g.Group(children),
	)
}