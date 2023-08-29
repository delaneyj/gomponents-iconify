package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myude(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Zm-7.46-8.07H39m-5.96-11.91H39m-5.96 5.95h3.89m-3.89-5.95v11.91"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.19 34.43V22.52h2.69a5.21 5.21 0 0 1 5.21 5.21v1.49a5.21 5.21 0 0 1-5.21 5.21Zm0-18.1v2.26a1.67 1.67 0 0 1-1.67 1.67h0a1.69 1.69 0 0 1-1.18-.49"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.19 13.58v2.75A1.66 1.66 0 0 1 19.52 18h0a1.67 1.67 0 0 1-1.67-1.67v-2.75M9 15.24a1.67 1.67 0 0 1 1.67-1.67h0a1.67 1.67 0 0 1 1.67 1.67v2.67M9 13.57v4.34"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.34 15.24A1.67 1.67 0 0 1 14 13.57h0a1.67 1.67 0 0 1 1.67 1.67v2.67M9 22.52v8a3.95 3.95 0 1 0 7.89 0v-8"/>`),
		g.Group(children),
	)
}