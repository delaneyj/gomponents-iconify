package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m41.98 18.143l-3.918-6.786l-10.144 5.857V5.5h-7.836v11.714L9.938 11.357L6.02 18.143L16.164 24L6.02 29.857l3.918 6.786l10.144-5.857V42.5h7.836V30.786l10.144 5.857l3.918-6.786L31.836 24l10.144-5.857z"/>`),
		g.Group(children),
	)
}