package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenScale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M5.021 5.001v5.938h7.957V5.001H5.021zm7.021 5.061H5.958V5.937h6.084v4.125z"/><path d="M7 7h3.917v1.938H7zm-6 4l1.387 1.375l-1.385 1.381l1.224 1.216l1.387-1.381L5 14.965V11H1zm14.763-9.96l-1.376 1.359L13 1.035v3.934h4l-1.387-1.365l1.375-1.358l-1.225-1.206zm-.002 13.931l1.226-1.215l-1.38-1.375L17 11h-4v3.965l1.381-1.368l1.38 1.374zM2.27 1.08L1.043 2.287l1.338 1.322L1 4.969h4V1.035l-1.393 1.37L2.27 1.08z"/></g>`),
		g.Group(children),
	)
}