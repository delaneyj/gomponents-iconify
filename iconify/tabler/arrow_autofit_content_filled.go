package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowAutofitContentFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M6.707 3.293a1 1 0 0 1 .083 1.32l-.083.094L5.415 6H10a1 1 0 0 1 .117 1.993L10 8H5.415l1.292 1.293a1 1 0 0 1 .083 1.32l-.083.094a1 1 0 0 1-1.32.083l-.094-.083l-3-3a1.008 1.008 0 0 1-.097-.112l-.071-.11l-.054-.114l-.035-.105l-.025-.118l-.007-.058L2 7l.003-.075l.017-.126l.03-.111l.044-.111l.052-.098l.064-.092l.083-.094l3-3a1 1 0 0 1 1.414 0zm11.906-.083l.094.083l3 3a.927.927 0 0 1 .097.112l.071.11l.054.114l.035.105l.03.148L22 7l-.003.075l-.017.126l-.03.111l-.044.111l-.052.098l-.074.104l-.073.082l-3 3a1 1 0 0 1-1.497-1.32l.083-.094L18.585 8H14a1 1 0 0 1-.117-1.993L14 6h4.585l-1.292-1.293a1 1 0 0 1-.083-1.32l.083-.094a1 1 0 0 1 1.32-.083zM18 13H6a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-2a3 3 0 0 0-3-3z"/></g>`),
		g.Group(children),
	)
}