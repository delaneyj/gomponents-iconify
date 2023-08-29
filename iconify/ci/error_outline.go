package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ErrorOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.99 22C6.468 21.994 1.996 17.515 2 11.993C2.004 6.472 6.482 1.998 12.003 2C17.525 2.002 22 6.478 22 12c-.003 5.525-4.485 10.002-10.01 10ZM4 12.172A8 8 0 1 0 4 12v.172ZM13 17h-2v-2h2v2Zm0-4h-2V7h2v6Z"/>`),
		g.Group(children),
	)
}