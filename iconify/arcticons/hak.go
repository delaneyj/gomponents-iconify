package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.866 18.753l-1.851 10.494m5.641 0L32.26 24l5.24-5.212M32.26 24h-1.32m-20.049 3.03l-.391 2.217m1.85-10.494l-.384 2.177m7.337-2.177l-1.851 10.494m9.222-3.51h-4.546m-1.747 3.479l5.255-10.463l1.56 10.494M14.905 23.98h3.476"/>`),
		g.Group(children),
	)
}