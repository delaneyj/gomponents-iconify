package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.195 24.003a6.748 6.748 0 0 0 6.749-6.748h0a6.748 6.748 0 0 0-6.749-6.749"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.195 37.5a6.748 6.748 0 0 0 6.749-6.748h0a6.748 6.748 0 0 0-6.749-6.749M15.06 35.223c1.863 1.561 3.876 2.277 8.395 2.277h2.74M15.056 12.755c1.868-1.557 3.883-2.267 8.402-2.255l2.74.006M19.32 24.003h6.875"/>`),
		g.Group(children),
	)
}