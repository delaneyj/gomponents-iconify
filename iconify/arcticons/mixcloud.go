package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mixcloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.014 24h-7.768M9.5 28.115v-8.239l4.124 8.248l4.123-8.235v8.235M38.5 20.392l-8.247 7.216m8.247 0l-8.247-7.216"/>`),
		g.Group(children),
	)
}