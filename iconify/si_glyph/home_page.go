package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomePage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m7.65 4.158l-4.485 4.94a.648.648 0 0 0 .002.849l2.868.005v5.024c0 .553.439 1 .982 1h1.965a.99.99 0 0 0 .982-1v-5.02l2.811.005a.65.65 0 0 0-.004-.848L8.414 4.159a.506.506 0 0 0-.764-.001zm6.268-2.324c0 .552-.437 1-.973 1H3.056c-.537 0-.973-.448-.973-1V1c0-.552.436-1 .973-1h9.889c.536 0 .973.448.973 1v.834z"/>`),
		g.Group(children),
	)
}