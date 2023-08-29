package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Male(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.875 5.884V.125h-5.762l2.223 2.221l-3.703 3.711a5.395 5.395 0 0 0-3.131-1.002c-3.004 0-5.447 2.452-5.447 5.467c0 3.016 2.443 5.468 5.447 5.468c3.002 0 5.444-2.452 5.444-5.468a5.45 5.45 0 0 0-1.011-3.161l3.702-3.712l2.238 2.235zM6.464 14.165a3.534 3.534 0 0 1-3.532-3.53a3.535 3.535 0 0 1 3.532-3.529a3.535 3.535 0 0 1 3.532 3.529a3.534 3.534 0 0 1-3.532 3.53z"/>`),
		g.Group(children),
	)
}