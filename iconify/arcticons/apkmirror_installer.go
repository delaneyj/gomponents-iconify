package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApkmirrorInstaller(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.999 34.944l-7.113-21.886L9.5 34.944m2.462-7.386h9.575M38.5 34.942l-7.113-21.886l-7.386 21.886"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}