package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AfasPocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm12.163 13.722h4.778M18.821 24h3.106m-2.264-4.778l-1.684 9.556"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.033 27.73a2.22 2.22 0 0 0 2.158 1.048h1.415a2.97 2.97 0 0 0 2.804-2.39h0A1.947 1.947 0 0 0 35.448 24h-1.563a1.947 1.947 0 0 1-1.962-2.389h0a2.97 2.97 0 0 1 2.804-2.389h1.415A2.22 2.22 0 0 1 38.3 20.27m-29 8.479l4.784-9.527m1.422 9.556l-1.422-9.556m.946 6.359h-4.139m11.503 3.168l4.784-9.527m1.422 9.556l-1.422-9.556m.946 6.359h-4.139"/>`),
		g.Group(children),
	)
}