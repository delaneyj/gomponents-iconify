package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Passman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c13.355-4.243 16.635-14.368 16.635-19.286V11.309a1.133 1.133 0 0 0-.739-1.062L24.788 4.64a2.266 2.266 0 0 0-1.576 0L8.104 10.247a1.133 1.133 0 0 0-.74 1.063v12.905c0 4.918 3.281 15.043 16.636 19.286Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.816 22.652l-2.81-3.822l-2.771 3.848L24 18.825l-4.516-1.447l4.519 1.44l-.02-4.741l.027 4.742l4.504-1.483l-4.502 1.49ZM18 30.1h12"/>`),
		g.Group(children),
	)
}