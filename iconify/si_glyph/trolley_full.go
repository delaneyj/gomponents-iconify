package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrolleyFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(0 2)"><path d="M5.8 9L3.451 0H.002l.025.938h2.438L5.065 10H15.97V9H5.8z"/><path d="M7 6h1.958v1.955H7zm3 0h1.953v2H10zm0-3h1.938v1.899H10zm0-3h1.953v1.941H10zm3 6h1.953v1.906H13zm0-3h1.938v1.899H13zm0-3h1.953v1.941H13zM7 3h1.938v1.899H7zm0-3h1.953v1.941H7z"/><ellipse cx="7.494" cy="10.471" rx="1.494" ry="1.471"/><circle cx="13.487" cy="10.487" r="1.487"/></g>`),
		g.Group(children),
	)
}