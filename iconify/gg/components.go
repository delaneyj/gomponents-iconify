package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Components(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.757 6.343L12 2.1l4.242 4.243L12 10.586L7.757 6.343Zm2.829 0L12 4.93l1.414 1.414L12 7.757l-1.414-1.414ZM2.1 12l4.243-4.243L10.586 12l-4.243 4.242L2.1 12Zm2.829 0l1.414-1.414L7.757 12l-1.414 1.414L4.93 12Zm8.485 0l4.243 4.242L21.899 12l-4.242-4.243L13.414 12Zm4.243-1.414L16.243 12l1.414 1.414L19.07 12l-1.414-1.414Zm-9.9 7.071L12 13.414l4.242 4.243L12 21.899l-4.243-4.242Zm2.829 0L12 16.243l1.414 1.414L12 19.07l-1.414-1.414Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}