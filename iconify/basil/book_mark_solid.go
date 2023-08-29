package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookMarkSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.02 3.5c.122 0 .215.108.199.229a20.944 20.944 0 0 0 .103 6.225l.014.087a.25.25 0 0 0 .456.095l1.624-2.501a.1.1 0 0 1 .168 0l1.624 2.501a.25.25 0 0 0 .456-.095l.014-.087c.345-2.06.379-4.158.103-6.225a.202.202 0 0 1 .2-.229H18.5A1.5 1.5 0 0 1 20 5v15a1 1 0 0 1-1 1H7.5a3.5 3.5 0 0 1-3.465-3H4V8a4.5 4.5 0 0 1 4.5-4.5h3.52Zm-4.52 12h11v4h-11a2 2 0 1 1 0-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}