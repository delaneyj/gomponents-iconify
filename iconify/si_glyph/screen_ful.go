package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenFul(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M4.021 5.001v5.938h7.957V5.001H4.021zm7.021 5.061H4.958V5.937h6.084v4.125z"/><path d="M6 7h3.917v1.938H6zm6-6l1.387 1.375l-1.385 1.381l1.224 1.216l1.387-1.381L16 4.965V1h-4zM2.763 11.04l-1.376 1.359L0 11.035v3.934h4l-1.387-1.365l1.375-1.358l-1.225-1.206zm-.002-6.069l1.226-1.215l-1.38-1.375L4 1H0v3.965l1.381-1.368l1.38 1.374zM13.27 11.08l-1.227 1.207l1.338 1.322L12 14.969h4v-3.934l-1.393 1.37l-1.337-1.325z"/></g>`),
		g.Group(children),
	)
}