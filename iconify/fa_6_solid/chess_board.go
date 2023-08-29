package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 32C28.7 32 0 60.7 0 96v320c0 35.3 28.7 64 64 64h320c35.3 0 64-28.7 64-64V96c0-35.3-28.7-64-64-64H64zm64 64v64h64V96h64v64h64V96h64v64h-64v64h64v64h-64v64h64v64h-64v-64h-64v64h-64v-64h-64v64H64v-64h64v-64H64v-64h64v-64H64V96h64zm64 128h64v-64h-64v64zm0 64v-64h-64v64h64zm64 0h-64v64h64v-64zm0 0h64v-64h-64v64z"/>`),
		g.Group(children),
	)
}