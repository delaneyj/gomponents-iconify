package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JtxBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-3.116 13.995l-7.799 10.334m7.799 0l-7.799-10.334"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.97 16.05v11.7c0 1.17.78 1.95 1.95 1.95h.585m-4.484-10.334h4.094m-7.87-.296v11.7a3.91 3.91 0 0 1-3.9 3.899h0c-.974 0-1.949-.39-2.729-1.17"/><circle cx="17.245" cy="14.391" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}