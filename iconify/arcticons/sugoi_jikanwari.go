package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SugoiJikanwari(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.3 37.5H14.7c-2.3 0-4.1-1.9-4.1-4.1V14.7c0-2.3 1.9-4.1 4.1-4.1h18.7c2.3 0 4.1 1.9 4.1 4.1v18.7c0 2.2-1.9 4.1-4.2 4.1ZM5.5 14.8h5m-5 9.2h5m-5 9.2h5M33.2 5.5v5m-9.2-5v5m-9.2-5v5m18.4 27v5m-9.2-5v5m-9.2-5v5m22.7-27.7h5m-5 9.2h5m-5 9.2h5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.031 17.5v8.031a4.969 4.969 0 1 0 9.938 0V17.5m-11.813 0h3.75m6.188 0h3.75"/>`),
		g.Group(children),
	)
}