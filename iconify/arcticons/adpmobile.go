package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adpmobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.56 19.478v5.087h-4.946"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.848 28.24h6.712v4.309h5.193a7.949 7.949 0 0 0 7.173-4.522h.775v4.522h4.027v-4.027h1.237a6.535 6.535 0 0 0 6.535-6.536h0a6.535 6.535 0 0 0-6.535-6.535H31.7v8.408a8.408 8.408 0 0 0-8.408-8.408h-8.831L4.5 32.549h4.981l2.367-4.31Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.951 28.24h1.413a4.38 4.38 0 0 0 4.38-4.381h0a4.38 4.38 0 0 0-4.38-4.38h-1.413m13.853 4.686h1.664a2.343 2.343 0 0 0 2.343-2.343h0a2.343 2.343 0 0 0-2.343-2.344h-1.664"/>`),
		g.Group(children),
	)
}