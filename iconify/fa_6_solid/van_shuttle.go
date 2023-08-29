package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VanShuttle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 104v88h96V96H72c-4.4 0-8 3.6-8 8zm482 88l-80.9-96H384v96h162zm-226 0V96h-96v96h96zm272 192h-16c0 53-43 96-96 96s-96-43-96-96H256c0 53-43 96-96 96s-96-43-96-96H48c-26.5 0-48-21.5-48-48V104c0-39.8 32.2-72 72-72h393.1c18.9 0 36.8 8.3 49 22.8L625 186.5c9.7 11.5 15 26.1 15 41.2V336c0 26.5-21.5 48-48 48zm-64 0a48 48 0 1 0-96 0a48 48 0 1 0 96 0zm-368 48a48 48 0 1 0 0-96a48 48 0 1 0 0 96z"/>`),
		g.Group(children),
	)
}