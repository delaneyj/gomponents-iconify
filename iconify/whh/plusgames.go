package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plusgames(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1006.54 241l-135 135q25 67 25 136t-25 136l134 134q18 18 18 44.5t-18 44.5l-134 134q-18 18-44.5 18t-44.5-18l-134-134q-67 25-136 25t-136-25l-135 135q-18 18-44 18t-45-18l-134-134q-18-19-18-45t18-44l135-135q-25-67-25-136t26-136l-135-135q-19-18-19-44.5t19-44.5l133-134q18-18 44.5-18t44.5 18l135 136q67-26 136-26t136 25l135-135q18-18 44-18t45 18l134 134q18 19 18 45t-18 44zm-493.5 79q-79.5 0-136 56.5T320.54 512t56.5 135.5t136 56.5t135.5-56.5t56-135.5t-56-135.5t-135.5-56.5z"/>`),
		g.Group(children),
	)
}