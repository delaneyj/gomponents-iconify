package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LrcEditor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 17.594v12.812h6.406m3.046 0V17.594h4.195a4.303 4.303 0 0 1 0 8.606h-4.195m4.195 0l4.194 4.203M38.5 26.11v.052a4.244 4.244 0 0 1-4.244 4.244h0a4.244 4.244 0 0 1-4.244-4.244v-4.324a4.244 4.244 0 0 1 4.244-4.244h0a4.244 4.244 0 0 1 4.244 4.244v.052"/>`),
		g.Group(children),
	)
}