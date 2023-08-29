package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaslowPyramids(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMaslowPyramids0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke-linejoin="round" d="m24 4l-9 15.98h18L24 4Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M24 19.98L24.008 44"/><path stroke-linecap="round" stroke-linejoin="round" d="M11.347 25.975L2 42h15.005"/><path stroke-linecap="round" d="M9.1 30.995h7.904"/><path stroke-linecap="round" stroke-linejoin="round" d="M36.748 25.975L46.094 42H31"/><path stroke-linecap="round" d="M39.094 30.995H31.1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMaslowPyramids0)"/>`),
		g.Group(children),
	)
}