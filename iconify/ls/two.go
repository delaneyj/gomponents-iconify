package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Two(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M423 368L156 682h325v72H0l255-299s117-135 128-152s19-47 19-75c0-87-70-158-158-158c-87 0-158 71-158 158c0 26 6 49 17 71H29c-8-23-13-45-13-71C16 103 119 0 244 0c117 0 213 88 226 202c1 8 1 18 1 26c0 53-18 102-48 140z"/>`),
		g.Group(children),
	)
}