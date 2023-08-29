package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YonoLiteSbi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.556 19.95v7.803h3.901m2.639-7.803v7.803m10.447 0h3.901m-3.901-7.803h3.901m-3.901 3.902h2.536m-2.536-3.902v7.803m-7.808-7.803h5.169m-2.536 7.803V19.95m-.254-3.098V9.049l5.17 7.803V9.049m-15.352 0l-2.536 3.901l-2.633-3.901m2.633 7.803V12.95m4.764 1.317a2.585 2.585 0 1 0 5.17 0v-2.633A2.623 2.623 0 0 0 19.697 9a2.542 2.542 0 0 0-2.536 2.634Zm15.908 0a2.585 2.585 0 1 0 5.17 0v-2.633A2.623 2.623 0 0 0 35.603 9a2.542 2.542 0 0 0-2.536 2.634Zm3.884 16.93V39m-15.284-.878a2.325 2.325 0 0 0 1.95.878h1.17a1.957 1.957 0 0 0 1.952-1.95h0a1.956 1.956 0 0 0-1.95-1.951h-1.27a1.957 1.957 0 0 1-1.95-1.951h0a1.957 1.957 0 0 1 1.95-1.951h1.17a2.094 2.094 0 0 1 1.951.878m5.839 3.024a1.95 1.95 0 1 1 0 3.901h-3.219v-7.803h3.219a1.95 1.95 0 1 1 0 3.901Zm0 0h-3.219m-13.242 3.754a3.968 3.968 0 1 0-2.006 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}