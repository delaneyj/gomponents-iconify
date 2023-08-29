package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadDoneRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.55 13.15L18 4.7q.3-.3.713-.3t.712.3q.3.3.3.712t-.3.713L10.25 15.3q-.3.3-.7.3t-.7-.3l-4.275-4.275q-.3-.3-.288-.713T4.6 9.6q.3-.3.713-.3t.712.3l3.525 3.55ZM6 20q-.425 0-.713-.288T5 19q0-.425.288-.713T6 18h12q.425 0 .713.288T19 19q0 .425-.288.713T18 20H6Z"/>`),
		g.Group(children),
	)
}