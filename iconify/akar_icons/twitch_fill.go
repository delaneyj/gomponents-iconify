package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwitchFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.547 1L1 4.776v15.433h5.5V23h3.093l2.922-2.791h4.47L23 14.462V1H2.547Zm18.39 12.478l-3.438 3.283H12l-2.922 2.791v-2.79h-4.64V2.97h16.499v10.508Zm-3.438-6.731v5.74h-2.062v-5.74H17.5Zm-5.499 0v5.74H9.938v-5.74H12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}