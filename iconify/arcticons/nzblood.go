package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nzblood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.368 4.5s-6.814 9.484-8.885 13.487S9.031 27.994 11.31 34.48c2.277 6.487 8.903 8.558 8.903 8.558s7.115 2.184 12.865-2.955c3.422-3.058 4.594-7.17 4.34-10.344c-.315-3.907-3.17-6.185-4.378-7.225c-2.734-2.355-4.957-4.251-7.049-8.39C23.703 9.596 23.368 4.5 23.368 4.5"/>`),
		g.Group(children),
	)
}