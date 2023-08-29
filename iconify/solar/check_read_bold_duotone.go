package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckReadBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M3.464 20.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535Z" opacity=".5"/><path d="M18.581 9.474a.75.75 0 1 0-1.162-.948l-5.168 6.33a.749.749 0 0 0-.879 1.116l.286.438a.75.75 0 0 0 1.209.064l5.714-7Zm-4 0a.75.75 0 1 0-1.162-.948l-5.133 6.288l-1.705-2.088a.75.75 0 0 0-1.162.948l2.286 2.8a.75.75 0 0 0 1.162 0l5.714-7Z"/></g>`),
		g.Group(children),
	)
}