package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCollectComputer0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M22 6H9a3 3 0 0 0-3 3v22a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3v-9M24 34v8m-10 0h20"/><path fill="#fff" d="M33.3 6C31.478 6 30 7.435 30 9.205c0 3.204 3.9 6.117 6 6.795c2.1-.678 6-3.59 6-6.795C42 7.435 40.523 6 38.7 6A3.326 3.326 0 0 0 36 7.362A3.326 3.326 0 0 0 33.3 6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCollectComputer0)"/>`),
		g.Group(children),
	)
}