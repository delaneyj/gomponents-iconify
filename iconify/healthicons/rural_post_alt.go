package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuralPostAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.172 21.102A4.858 4.858 0 0 0 8 26.081V42h4V29h8v13h20V26.077a4.858 4.858 0 0 0 1.823-4.976l-1.113-4.875a1.517 1.517 0 0 0-1.479-1.179H31v-5a4 4 0 0 0-4-4h-6a4 4 0 0 0-4 4v5H8.764c-.707 0-1.321.49-1.479 1.18l-1.113 4.874Zm1.95.445l1.027-4.499h4.627v5.137a2.863 2.863 0 1 1-5.654-.637Zm7.827-4.499h1.177c.444 1.725 2.01 3 3.874 3h1.911v1.52a3.481 3.481 0 0 1-6.962 0v-4.52Zm9.05-7.5a1 1 0 1 0-2 0v2.5h-2.5a1 1 0 1 0 0 2H23v2.5a1 1 0 1 0 2 0v-2.5h2.5a1 1 0 1 0 0-2H25v-2.5Zm.085 12.02v-1.52H27a4.002 4.002 0 0 0 3.874-3h1.172v4.52a3.481 3.481 0 0 1-6.962 0Zm9.135-4.52h4.627l1.028 4.5a2.863 2.863 0 1 1-5.655.637v-5.137Zm-10.22 12h13v10H24v-10Zm11 2h-9v6h9v-6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}