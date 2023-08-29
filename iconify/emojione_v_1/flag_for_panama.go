package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForPanama(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#2b3990" d="M0 32v11c0 6.075 3.373 11 10 11h22V32H0z"/><path fill="#e6e7e8" d="M32 10H10C3.373 10 0 14.925 0 21v11h32V10zm0 44h22c6.627 0 10-4.925 10-11V32H32v22z"/><path fill="#ec1c24" d="M54 10H32v22h32V21c0-6.075-3.373-11-10-11"/><path fill="#cb0000" d="m54.24 41.527l-4.01.008l-1.241-4.056l-1.236 4.056l-4.01-.008L46.991 44l-1.254 4.03l3.26-2.505l3.265 2.505l-1.26-4.03z"/><path fill="#2b3990" d="m19.993 19.778l-4.01.007l-1.244-4.056l-1.236 4.056l-4.01-.007l3.248 2.472l-1.254 4.03l3.26-2.504l3.265 2.504l-1.26-4.03z"/>`),
		g.Group(children),
	)
}