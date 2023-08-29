package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionsBookmarkOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 16h12V4h-2v6.125q0 .3-.25.45t-.5-.025l-1.225-.725q-.25-.15-.537-.15t-.513.15l-1.225.725q-.275.175-.513.025t-.237-.45V4H8v12Zm0 2q-.825 0-1.413-.588T6 16V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H8Zm-4 4q-.825 0-1.413-.588T2 20V7q0-.425.288-.713T3 6q.425 0 .713.288T4 7v13h13q.425 0 .713.288T18 21q0 .425-.288.713T17 22H4ZM8 4v12V4Zm5 6.125q0 .3.238.45t.512-.025l1.225-.725q.225-.15.513-.15t.537.15l1.225.725q.25.175.5.025t.25-.45q0 .3-.25.45t-.5-.025l-1.225-.725q-.25-.15-.537-.15t-.513.15l-1.225.725q-.275.175-.513.025t-.237-.45Z"/>`),
		g.Group(children),
	)
}