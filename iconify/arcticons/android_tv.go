package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AndroidTv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.367 35.5H41.5a2 2 0 0 0 2-2v-24a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h9.132"/><circle cx="24" cy="31" r="9.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="31" r="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.937 24.481L24 23.85l-.937.631m7.456 7.456L31.15 31l-.631-.937m-7.456 7.456l.937.631l.937-.631m-7.456-7.456L16.85 31l.631.937"/>`),
		g.Group(children),
	)
}