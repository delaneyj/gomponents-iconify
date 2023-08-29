package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WriteJapanese(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM14.263 19.299L5.5 33.609M31.35 7.865L14.262 19.3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 22.627H28.217c-3.207.036-7.735 3.124-7.708 7.866V42.5M31.35 7.865l-1.623 14.762M26.42 11.164l4.275 2.654m-16.432 5.481l8.686 5.681"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.515 25.805c-.775 4.437-.584 8.46.282 12.328"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.042 27.957c3.475.354 6.918.165 10.333-.473m-1.825 2.107c-.183 5.724-8.751 11.11-8.666 6.469c.044-2.413 3.273-5.156 6.513-5.195c2.47-.03 5.367.883 5.387 3.099c.026 2.933-2.37 4.625-4.936 4.834"/>`),
		g.Group(children),
	)
}