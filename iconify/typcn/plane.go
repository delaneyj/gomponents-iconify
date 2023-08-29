package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.996 13.507L14 10.081V4.125c0-.827-.673-1.5-1.5-1.5s-1.5.673-1.5 1.5v5.956l-5.996 3.426a1 1 0 0 0 .77 1.829L11 13.844v4.451l-1.625 1.3a1 1 0 0 0 .996 1.709l2.129-.852l2.129.852a.998.998 0 0 0 1.235-.426a1.001 1.001 0 0 0-.239-1.284L14 18.294v-4.451l5.226 1.493l.274.039a1 1 0 0 0 .496-1.868zM12.5 4.375a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1z"/>`),
		g.Group(children),
	)
}