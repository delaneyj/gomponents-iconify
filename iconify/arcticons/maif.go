package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maif(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" d="M32.371 8.714L4.5 33.806l39 5.48L32.371 8.714ZM31.859 24v6.638"/><path stroke-linecap="round" stroke-linejoin="round" d="M15.392 30.631V24l3.319 6.638l3.319-6.628v6.628m12.586-3.319h2.158m-2.158 3.319V24h3.319"/><path d="M28.38 28.418h-2.876"/><path stroke-linecap="round" stroke-linejoin="round" d="M24.787 30.618L26.945 24l2.157 6.638"/></g>`),
		g.Group(children),
	)
}