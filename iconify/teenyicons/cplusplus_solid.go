package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CplusplusSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5.421L14 4.213v6.574L7.5 14.58L1 10.787V4.213L7.5.42ZM6.5 4a3.5 3.5 0 1 0 0 7h.586a3.914 3.914 0 0 0 2.768-1.146l-.708-.708a2.914 2.914 0 0 1-2.06.854H6.5a2.5 2.5 0 0 1 0-5h.586a2.91 2.91 0 0 1 2.06.854l.708-.708A3.914 3.914 0 0 0 7.086 4H6.5ZM7 7V6h1v1h2V6h1v1h1v1h-1v1h-1V8H8v1H7V8H6V7h1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}