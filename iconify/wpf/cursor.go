package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cursor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="m7 2.014l13.162 12.328l-5.377.488l-.977.088l.406.894l3.263 7.145l-2.406 1.058l-3.113-7.222l-.391-.91l-.722.678l-3.819 3.582L7 2.014m0-2a2 2 0 0 0-2 2.003l.026 18.127a2.003 2.003 0 0 0 1.999 1.998c.5 0 .991-.188 1.369-.541l2.463-2.311l2.377 5.516a2.006 2.006 0 0 0 1.836 1.208c.274 0 .549-.057.806-.169l2.406-1.059a1.998 1.998 0 0 0 1.013-2.661l-2.498-5.47l3.544-.322a1.998 1.998 0 0 0 1.187-3.451L8.367.554A1.999 1.999 0 0 0 7 .014z"/>`),
		g.Group(children),
	)
}