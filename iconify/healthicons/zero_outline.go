package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZeroOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 18a9 9 0 0 1 9-9h4c2.303 0 4.41.87 6 2.292A8.985 8.985 0 0 1 35 18v12a9 9 0 0 1-9 9h-4a8.985 8.985 0 0 1-6.708-3A8.977 8.977 0 0 1 13 30V18Zm9-7a7 7 0 0 0-7 7v12c0 1.792.675 3.43 1.783 4.667A6.985 6.985 0 0 0 22 37h4a7 7 0 0 0 7-7V18a6.985 6.985 0 0 0-2.333-5.217A6.977 6.977 0 0 0 26 11h-4Zm0 4a3 3 0 0 0-3 3v7.817L25.59 15H22Zm-5 3a5 5 0 0 1 5-5h4a5 5 0 0 1 1.561.248a1 1 0 0 1 .543 1.47l-9.25 15.183a1 1 0 0 1-1.854-.52V18Zm13.163-1.543a1 1 0 0 1 .788.84c.032.23.049.465.049.703v12a5 5 0 0 1-5 5h-4a4.98 4.98 0 0 1-2.318-.568a1 1 0 0 1-.39-1.406l9.815-16.11a1 1 0 0 1 1.056-.459Zm-8.502 16.524c.111.013.224.019.339.019h4a3 3 0 0 0 3-3v-9.065l-7.339 12.046Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}