package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DomainVerificationRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.95 13.65l-1.375-1.375q-.3-.3-.725-.3t-.725.3q-.3.3-.3.725t.3.725l2.125 2.125q.3.3.7.3t.7-.3l4.225-4.225q.3-.3.3-.725t-.3-.725q-.3-.3-.725-.3t-.725.3L10.95 13.65ZM4 8h16V6H4v2Zm0 12q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Z"/>`),
		g.Group(children),
	)
}