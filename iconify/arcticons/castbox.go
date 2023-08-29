package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Castbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.115 3L6.019 13.791v20.522L24.115 45l17.866-10.687V13.791L24.114 3.001ZM13.394 21.243v5.514v-5.514Zm.023 2.748h4.93h-4.93Zm9.695-7.452V28.25V16.54Zm4.678 1.083V28.35V17.622Zm-9.42 1.222v12.618v-12.617Zm.029 3.243h4.63h-4.63Zm4.798 1.126h4.534h-4.535Zm14.263-.801v3.176m-5.006-5.07v6.963"/>`),
		g.Group(children),
	)
}