package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 0C28.7 0 0 28.7 0 64v384c0 35.3 28.7 64 64 64h256c35.3 0 64-28.7 64-64V64c0-35.3-28.7-64-64-64H64zm32 64h192c17.7 0 32 14.3 32 32v32c0 17.7-14.3 32-32 32H96c-17.7 0-32-14.3-32-32V96c0-17.7 14.3-32 32-32zm32 160a32 32 0 1 1-64 0a32 32 0 1 1 64 0zM96 352a32 32 0 1 1 0-64a32 32 0 1 1 0 64zm-32 64c0-17.7 14.3-32 32-32h96c17.7 0 32 14.3 32 32s-14.3 32-32 32H96c-17.7 0-32-14.3-32-32zm128-160a32 32 0 1 1 0-64a32 32 0 1 1 0 64zm32 64a32 32 0 1 1-64 0a32 32 0 1 1 64 0zm64-64a32 32 0 1 1 0-64a32 32 0 1 1 0 64zm32 64a32 32 0 1 1-64 0a32 32 0 1 1 64 0zm-32 128a32 32 0 1 1 0-64a32 32 0 1 1 0 64z"/>`),
		g.Group(children),
	)
}