package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M31.22 22.044h7.28l-7.28 9.645h7.28M27.408 19.041v12.648m-1.911-9.645h3.822"/><circle cx="22.321" cy="17.585" r="1.274"/><path d="M22.321 22.044v9.645m-6.815-7.279a3.64 3.64 0 0 1 0 7.28H9.5V17.13h6.006a3.64 3.64 0 0 1 0 7.28Zm0 0H9.5"/></g>`),
		g.Group(children),
	)
}