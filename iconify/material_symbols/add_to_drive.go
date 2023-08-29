package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddToDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 21v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2Zm-12.375.5q-.525 0-.975-.275t-.725-.725l-2.35-4.1q-.275-.45-.275-1t.275-1l6.35-10.9q.275-.5.738-.75t.987-.25h4.7q.525 0 .988.25t.737.75l4.475 7.675q-.375-.125-.75-.162t-.8-.038h-.425q-.2 0-.425.05L14.35 4.5h-4.7L3.3 15.4l2.35 4.1h7.9q.275.575.638 1.075t.862.925H5.625ZM7.25 17l-.725-1.275L11.1 7.75h1.8l2.525 4.4q-.45.35-.787.712t-.638.813L12 10.2L9.25 15h4.1q-.2.5-.275.988T13 17H7.25Z"/>`),
		g.Group(children),
	)
}