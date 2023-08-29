package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DomainVerificationOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h16V8H4v10Zm6.95-4.35l3.475-3.475q.3-.3.725-.3t.725.3q.3.3.3.725t-.3.725L11.65 15.85q-.3.3-.7.3t-.7-.3l-2.125-2.125q-.3-.3-.3-.725t.3-.725q.3-.3.725-.3t.725.3l1.375 1.375ZM4 18V6v12Z"/>`),
		g.Group(children),
	)
}