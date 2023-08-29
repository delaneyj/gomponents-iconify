package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Insertpicturecenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992.59 192h-960q-13 0-22.5-9.5T.59 160t9.5-22.5t22.5-9.5h960q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5zm-192-128h-768q-13 0-22.5-9.5T.59 32t9.5-22.5T32.59 0h768q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5zm-480 640q-26 0-45-18.5t-19-45.5V320q0-26 19-45t45-19h384q27 0 45.5 19t18.5 45v320q0 27-18.5 45.5t-45.5 18.5h-384zm-288 64h896q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-896q-13 0-22.5-9.5T.59 800t9.5-22.5t22.5-9.5zm0 128h768q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-768q-13 0-22.5-9.5T.59 928t9.5-22.5t22.5-9.5z"/>`),
		g.Group(children),
	)
}