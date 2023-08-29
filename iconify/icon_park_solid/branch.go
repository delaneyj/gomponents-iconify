package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Branch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBranch0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M40 28a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM9 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 32a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke-linecap="round" d="M9 12v24v-11.992h27"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBranch0)"/>`),
		g.Group(children),
	)
}