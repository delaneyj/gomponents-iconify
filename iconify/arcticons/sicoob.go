package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sicoob(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.239 41.163L5.127 9.848a.945.945 0 0 1 .815-1.42l36.113-.063a.945.945 0 0 1 .82 1.416l-18.002 31.38a.942.942 0 0 1-1.633.003Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.369 22.074l-8.742.083l4.3-7.632Zm-4.441-7.557L20.463 8.47m-.823 13.685l-3.638 6.307m12.372-6.384h7.16"/>`),
		g.Group(children),
	)
}