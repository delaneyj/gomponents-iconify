package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DischargeLounge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 6a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H9Zm14 5v3h-3v2h3v3h2v-3h3v-2h-3v-3h-2ZM10.04 21.275a1 1 0 0 1 1.922-.55l1.2 4.198H21v2h-1.308v1.692h.539a1 1 0 0 1 .39 1.921l1.331 4.16A1 1 0 0 1 22 35h4c0-.1.015-.203.047-.305l1.331-4.159a1 1 0 0 1 .391-1.92h.539v-1.693H27v-2h7.84l1.198-4.198a1 1 0 0 1 1.924.55l-1.935 6.771l1.289 6.767a.993.993 0 0 1 .017.187H39a1 1 0 1 1 0 2H9a1 1 0 1 1 0-2h1.666c0-.062.006-.124.018-.187l1.29-6.767l-1.935-6.771Zm24.06 7.34h-3.791v-1.692h3.96l-.23.802a1 1 0 0 0-.02.462l.081.428ZM34.553 31H29.33l-1.28 4h7.265l-.762-4Zm-16.86-2.385v-1.692h-3.96l.228.802a1 1 0 0 1 .021.462l-.081.428h3.791ZM12.684 35l.762-4h5.223l1.28 4h-7.265Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}