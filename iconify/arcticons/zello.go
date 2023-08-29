package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zello(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.31 14.31v19.38h10.57v8.81H8.144c-1.506-.292-2.346-1.211-2.643-2.643V25.762c0-1.762.881-3.524 1.762-4.405l7.048-7.047Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 14.31h28.19l-7.047-7.048C25.762 6.38 24 5.5 22.238 5.5H8.143c-1.64.302-2.33 1.33-2.643 2.643v6.167ZM24.881 42.5l8.81-8.81H42.5v6.167c-.301 1.392-1.094 2.35-2.643 2.643H24.881Zm8.809-37v28.19l7.048-7.047C41.62 25.762 42.5 24 42.5 22.238V8.143c-.296-1.596-1.292-2.336-2.643-2.643H33.69Z"/>`),
		g.Group(children),
	)
}