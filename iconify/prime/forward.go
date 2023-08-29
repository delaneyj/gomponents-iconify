package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.51 11.45l-7.47-7a.756.756 0 0 0-.81-.14c-.27.12-.45.39-.45.69v6.27L4.51 4.45a.756.756 0 0 0-.81-.14c-.27.12-.45.39-.45.69v14c0 .3.18.57.45.69c.1.04.2.06.3.06c.19 0 .37-.07.51-.2l7.27-6.82V19c0 .3.18.57.45.69c.1.04.2.06.3.06c.19 0 .37-.07.51-.2l7.47-7c.15-.14.24-.34.24-.55s-.09-.41-.24-.55ZM4.75 17.27V6.73L10.37 12l-5.62 5.27Zm8.53 0V6.73L18.9 12l-5.62 5.27Z"/>`),
		g.Group(children),
	)
}