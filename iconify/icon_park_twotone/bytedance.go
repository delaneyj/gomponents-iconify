package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bytedance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBytedance0"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="m5 7l5 2v28l-5 2V7Zm11 16l5 2v12l-5 2V23Zm11-2l5-2v16l-5-2V21ZM38 9l5 2v26l-5 2V9Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBytedance0)"/>`),
		g.Group(children),
	)
}