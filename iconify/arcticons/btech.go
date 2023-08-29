package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Btech(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.585 24.019c1.012 0 1.84.828 1.84 1.84s-.828 1.839-1.84 1.839H9.55V20.34h3.035c1.012 0 1.84.828 1.84 1.84s-.828 1.839-1.84 1.839Zm0 0H9.55m6.384-3.679h4.782m-2.391 7.358V20.34m4.044 3.679h2.391m1.287 3.679h-3.678V20.34h3.678m6.005 4.74h0a2.473 2.473 0 0 1-2.483 2.483h0c-1.38 0-2.484-1.104-2.484-2.392v-2.483c0-1.38 1.104-2.483 2.484-2.391h0a2.375 2.375 0 0 1 2.39 2.39h0m1.716-2.347v7.358m4.875-7.358v7.358m-4.875-3.771h4.875"/><circle cx="16.049" cy="27.59" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}