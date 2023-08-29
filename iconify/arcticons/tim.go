package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.382 4.5h19.235a6.648 6.648 0 0 1 6.663 6.662V29.83a6.648 6.648 0 0 1-6.662 6.662h-3.39V43.5l-8.63-7.007h-7.215A6.648 6.648 0 0 1 7.72 29.83V11.162A6.648 6.648 0 0 1 14.383 4.5Zm9.546 5.875V21.28m10.335-3.473L23.929 21.28m6.487 8.368L23.93 21.28m-6.697 8.33l6.695-8.33M13.24 17.665l10.688 3.615"/>`),
		g.Group(children),
	)
}