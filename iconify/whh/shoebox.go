package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shoebox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992.268 192h-960q-13 0-22.5-9.5T.268 160V32q0-13 9.5-22.5t22.5-9.5h960q13 0 22.5 9.5t9.5 22.5v128q0 13-9.5 22.5t-22.5 9.5zm-896 64h832q13 0 22.5 9.5t9.5 22.5v448q0 13-9.5 22.5t-22.5 9.5h-832q-13 0-22.5-9.5t-9.5-22.5V288q0-13 9.5-22.5t22.5-9.5zm224 352q0 13 9.5 22.5t22.5 9.5h320q13 0 22.5-9.5t9.5-22.5V416q0-13-9.5-22.5t-22.5-9.5h-320q-13 0-22.5 9.5t-9.5 22.5v192z"/>`),
		g.Group(children),
	)
}