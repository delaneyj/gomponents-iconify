package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M370 183q19 0 33 13t13.5 32.5T403 261t-33 13.5t-33-13.5t-13-33t13-32t33-13zm0-163q87 0 153 27t110 62t71 68t33 46q3 6 0 12q-5 8-18 26t-34 41t-52 47t-69 43t-87 32t-107 13t-106-13t-88-32t-69-43t-51-47t-35-40t-18-27q-5-6 0-12q7-13 33-46t71-68t111-62t152-27zm0 325q24 0 45-9t37-25t25-37t9-45t-9-45t-25-37t-37-25t-45-9t-45 9t-37 25t-25 37t-9 45t9 45t25 37t37 25t45 9z"/>`),
		g.Group(children),
	)
}