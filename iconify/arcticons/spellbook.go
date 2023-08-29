package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spellbook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.4 6.45v35.1a2 2 0 0 0 1.95 2h2.38V4.5h-2.38A2 2 0 0 0 8.4 6.45Zm4.33-1.95v39h24.92a2 2 0 0 0 2-2V6.45a2 2 0 0 0-2-1.95Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.16 12.45l10 5.78v11.54l-10 5.78l-10-5.78V18.23l10-5.78z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.16 17.59l-6.07 10.52h12.15l-6.08-10.52zm6.08 10.52l3.92-9.88l-10-.64l6.08 10.52z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.16 17.59l-10 .64l3.93 9.88l6.07-10.52z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.09 28.11l6.07 7.44l-10-5.78l3.93-1.66zm12.15 0l-6.08 7.44l10-5.78l-3.92-1.66zm-6.08-15.66v5.14"/>`),
		g.Group(children),
	)
}