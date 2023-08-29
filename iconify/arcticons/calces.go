package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calces(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.855 8.5h32.29v16.194H7.855zm0 20.645h7.73v3.865h-7.73zm24.56 0h7.73v3.865h-7.73z"/><circle cx="24" cy="31.077" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.855 36.907h5v3h-5zm9.097 0h5v2.999h-5zm9.096 0h5v3h-5zm9.097 0h5v3h-5zM9.75 21.67c0 .515.587.891 1.292.891c.67 0 1.094-.72 1.246-2.48l.224-2.673c.106-1.3.529-4.328.67-4.918c.2-.816.658-1.858 1.775-1.858c.705 0 1.293.376 1.293.892c0 .322-.195.675-.43.675m21.381 4.374l-3.099 4.106m3.099 0l-3.099-4.106m-9.072 0l-3.099 4.106m3.099 0l-3.099-4.106m-1.014-1.875a6.711 6.711 0 0 0-1.55 3.932a6.712 6.712 0 0 0 1.55 3.932m5.128-7.864a6.711 6.711 0 0 1 1.55 3.932a6.712 6.712 0 0 1-1.55 3.932m6.225-4.44a1.55 1.55 0 0 0-1.55-1.55h0a1.55 1.55 0 0 0-1.549 1.55v1.008a1.55 1.55 0 0 0 1.55 1.55h0a1.55 1.55 0 0 0 1.549-1.55m0 1.549v-6.198M16.23 20.68v-5.114a1.085 1.085 0 0 1 1.085-1.085h0a1.263 1.263 0 0 1 1.095.454m-3.099 1.638h2.169"/>`),
		g.Group(children),
	)
}