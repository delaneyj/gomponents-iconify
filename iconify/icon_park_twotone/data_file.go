package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDataFile0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M8 44V4h23l9 10.5V44H8Z"/><path fill="#555" d="M15 28h6v7h-6z"/><path d="M14 35h20"/><path fill="#555" d="M21 23h6v12h-6zm6-5h6v17h-6z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDataFile0)"/>`),
		g.Group(children),
	)
}