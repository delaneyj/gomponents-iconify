package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GamepadMinimalisticBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M8.5 11v3M7 12.5h3m-7.377 2.618l-.543 2.259c-.381 1.587.633 3.175 2.264 3.546c1.376.312 2.791-.342 3.413-1.576l.128-.255c.57-1.133 1.754-1.852 3.051-1.852h2.128c1.297 0 2.48.719 3.051 1.852l.128.255c.622 1.234 2.037 1.888 3.413 1.575c1.631-.37 2.645-1.958 2.264-3.545l-1.085-4.517c-.613-2.553-.92-3.83-1.79-4.666c-.22-.212-.463-.4-.725-.56C17.284 7 15.937 7 13.243 7h-2.486c-2.694 0-4.041 0-5.077.634c-.262.16-.505.348-.725.56c-.627.603-.961 1.434-1.328 2.806M12 7V6a1 1 0 0 1 1-1h1a1 1 0 0 0 1-1V3"/><path fill="currentColor" d="M16 11a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm2 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g>`),
		g.Group(children),
	)
}