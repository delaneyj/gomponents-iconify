package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableCells(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 32C28.7 32 0 60.7 0 96v320c0 35.3 28.7 64 64 64h384c35.3 0 64-28.7 64-64V96c0-35.3-28.7-64-64-64H64zm88 64v64H64V96h88zm56 0h88v64h-88V96zm240 0v64h-88V96h88zM64 224h88v64H64v-64zm232 0v64h-88v-64h88zm64 0h88v64h-88v-64zM152 352v64H64v-64h88zm56 0h88v64h-88v-64zm240 0v64h-88v-64h88z"/>`),
		g.Group(children),
	)
}