package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BagRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.494 4.011a.49.49 0 0 1-.489-.492a2.512 2.512 0 0 0-2.503-2.514A2.511 2.511 0 0 0 4 3.519a.49.49 0 1 1-.979 0C3.021 1.59 4.583.021 6.502.021c1.92 0 3.482 1.569 3.482 3.498a.49.49 0 0 1-.49.492z"/><path d="M12.492 4.011a.49.49 0 0 1-.488-.492c0-1.386-1.119-2.514-2.492-2.514a.49.49 0 0 1-.488-.492a.49.49 0 0 1 .488-.492c1.912 0 3.469 1.569 3.469 3.498a.491.491 0 0 1-.489.492zM11 14h5v1h-5zm3.371-10.939h-1.329v9.897h1.923V3.626c0-.313-.267-.565-.594-.565z"/><path d="M11.951 13.047L11.956 3H1.644a.614.614 0 0 0-.601.6v11.773c0 .312.227.627.557.627H10v-2.953h1.951z"/></g>`),
		g.Group(children),
	)
}