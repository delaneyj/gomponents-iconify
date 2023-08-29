package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Etlab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.835 31.816a18.988 18.988 0 0 1-.127-2.2c0-8.879 6.16-16.32 14.438-18.281a16.236 16.236 0 0 0-8.662-2.488c-9.01 0-16.312 7.304-16.312 16.313s7.303 16.312 16.312 16.312c5.263 0 9.944-2.493 12.927-6.363h0C38.645 41.337 31.809 45.5 24 45.5C12.126 45.5 2.5 35.875 2.5 24S12.126 2.5 24 2.5a21.4 21.4 0 0 1 11.578 3.38"/>`),
		g.Group(children),
	)
}