package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Money(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M432 64H16v320h416Zm-32 288H48V96h352Z"/><path fill="currentColor" d="M464 144v272H96v32h400V144h-32z"/><path fill="currentColor" d="M224 302.46c39.7 0 72-35.137 72-78.326s-32.3-78.326-72-78.326s-72 35.136-72 78.326s32.3 78.326 72 78.326Zm0-124.652c22.056 0 40 20.782 40 46.326s-17.944 46.326-40 46.326s-40-20.782-40-46.326s17.944-46.326 40-46.326ZM80 136h32v176H80zm256 0h32v176h-32z"/>`),
		g.Group(children),
	)
}