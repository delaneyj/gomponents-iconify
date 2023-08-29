package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LicenseSm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 6.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM10 4a3.997 3.997 0 0 1-1.842 3.369l.517 2.843L9 12l-1.66-.741L6 10.66l-1.34.599L3 12l.325-1.788l.517-2.843A4 4 0 1 1 10 4ZM5.252 8h1.496l.268 1.47l-.404-.18L6 9.017l-.612.273l-.404.18L5.252 8ZM6 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}