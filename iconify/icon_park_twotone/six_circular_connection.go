package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SixCircularConnection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSixCircularConnection0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M37.856 20v8M27.464 38l3.464-2l3.464-2m-13.856 4l-3.465-2l-3.464-2m-3.463-14v8m3.463-14l3.465-2l3.464-2m6.928 0l3.464 2l3.464 2"/><path fill="#555" d="M24 44a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0-32a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm14-8a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM10 20a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSixCircularConnection0)"/>`),
		g.Group(children),
	)
}