package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picnic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.635 22.35V10.5h3.88a3.98 3.98 0 0 1 0 7.96h-3.88m0 19.04V25.65l7.85 11.85V25.65"/><circle cx="24" cy="24.005" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 27.446V37.5m0-27v10.054m12.365-2.178v.048a3.925 3.925 0 0 1-3.925 3.926h0a3.925 3.925 0 0 1-3.925-3.925v-4A3.925 3.925 0 0 1 32.44 10.5h0a3.925 3.925 0 0 1 3.925 3.925v.049m0 19.052v.049A3.925 3.925 0 0 1 32.44 37.5h0a3.925 3.925 0 0 1-3.925-3.925v-4a3.925 3.925 0 0 1 3.925-3.924h0a3.925 3.925 0 0 1 3.925 3.925v.048"/>`),
		g.Group(children),
	)
}