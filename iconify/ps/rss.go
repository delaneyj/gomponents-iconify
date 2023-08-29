package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M106 358q18 18 18 43t-18 43t-43 18t-43-18Q2 427 2 401q0-27 18-43q18-18 43-18t43 18zm0 0zm-86 86zM2 158v88q88 0 152 64q63 61 63 152h89q0-126-90-214q-88-90-214-90zM2 2v88q154 0 263 109t109 263h88q0-94-35.5-178T327 137q-63-64-147-99.5T2 2z"/>`),
		g.Group(children),
	)
}