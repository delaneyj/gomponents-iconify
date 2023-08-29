package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Websocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#231F20" d="M192.44 144.645h31.78V68.339l-35.805-35.804l-22.472 22.472l26.497 26.497v63.14Zm31.864 15.931H113.452L86.954 134.08l11.237-11.236l21.885 21.885h45.028l-44.357-44.441l11.32-11.32l44.357 44.358v-45.03l-21.801-21.801l11.152-11.153L110.685 0H0l31.696 31.696v.084h65.74l23.227 23.227l-33.96 33.96L63.476 65.74V47.712h-31.78v31.193l55.007 55.007L64.314 156.3l35.805 35.805H256l-31.696-31.529Z"/>`),
		g.Group(children),
	)
}