package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shaadi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.327 39.078a9.036 9.036 0 0 0 6.356 4.363c2.328.132 5.444 0 7.258 0c3.632 0 8.05-4.363 8.05-9.696a8.731 8.731 0 0 0-7.29-8.762s-4.703-.962-8.878-1.794c-3.818-.762-5.813-3.504-5.813-8.837c0-9.115 4.63-9.412 6.355-9.696a11.88 11.88 0 0 1 3.814 0a7.756 7.756 0 0 1 6.006 4.538"/>`),
		g.Group(children),
	)
}