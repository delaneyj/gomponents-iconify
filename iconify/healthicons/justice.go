package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Justice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M25 6h-2v4.17a3.004 3.004 0 0 0-1.97 2.409l-7.29 1.955a1 1 0 0 0-.557.39l-5.855 8.294H6.12a8.52 8.52 0 0 0 .337 2.379c.95 3.256 3.797 5.62 7.163 5.62c3.495 0 6.431-2.55 7.264-6a8.525 8.525 0 0 0 .236-1.996v-.003h-.935l-4.6-7.108l5.842-1.567A3.01 3.01 0 0 0 23 15.829V38h-3v2h-6v2h20v-2h-6v-2h-3V15.83a3.002 3.002 0 0 0 2-2.783l5.14-1.378L28.473 17H27c0 .69.082 1.36.236 2c.833 3.45 3.77 6 7.264 6c3.365 0 6.213-2.364 7.163-5.621A8.499 8.499 0 0 0 42 17.003V17h-1.946L35.36 9.49a1 1 0 0 0-1.12-.456l-7.88 2.114a3.006 3.006 0 0 0-1.36-.977V6Zm5.901 11h6.795l-3.236-5.177L30.901 17Zm-13.098 6.218l-3.839-5.933l-4.188 5.933h8.027Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}