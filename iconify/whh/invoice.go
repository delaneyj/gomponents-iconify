package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Invoice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m832.417 960l-64 64l-64-64l-64 64l-64-64l-64 64l-64-64l-64 64l-64-64l-64 64l-64-64l-64 64l-64-64l-64 64V64q0-26 19-45t45-19h768q26 0 45 19t19 45v960zm-672-128h192q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5h-192q-13 0-22.5 9.5t-9.5 22.5t9.5 22.5t22.5 9.5zm608-672q0-13-9.5-22.5t-22.5-9.5h-576q-13 0-22.5 9.5t-9.5 22.5v384q0 13 9.5 22.5t22.5 9.5h352v96q0 13 9.5 22.5t22.5 9.5h192q13 0 22.5-9.5t9.5-22.5V160zm-192 416h128v64h-128v-64zm0-128h128v64h-128v-64zm0-128h128v64h-128v-64zm0-128h128v64h-128v-64zm-384 256h320v64h-320v-64zm0-128h320v64h-320v-64zm0-128h320v64h-320v-64z"/>`),
		g.Group(children),
	)
}