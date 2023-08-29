package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.523 5.5H7.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-4.948"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.538 3.367a9.554 9.554 0 0 0-9.554 9.554c0 7.478 7.313 16.49 9.207 18.7a.577.577 0 0 0 .881-.005c1.862-2.216 9.02-11.222 9.02-18.694a9.554 9.554 0 0 0-9.554-9.555Zm0 13.125a3.57 3.57 0 1 1 3.57-3.57a3.57 3.57 0 0 1-3.57 3.57Zm5.252 8.566l7.71-1.188m-37 5.702l19.82-3.055m-7.412 1.143L14.168 5.5"/>`),
		g.Group(children),
	)
}