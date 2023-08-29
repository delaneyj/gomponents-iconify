package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialDistancingProtectShieldTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M.75 9.25v4m22.5-4v4m-22.5-2h5.5m0-4.013v4.49a7.022 7.022 0 0 0 4.5 6.554l.614.236c.409.157.861.157 1.27 0l.614-.236a7.022 7.022 0 0 0 4.5-6.554v-4.49a.877.877 0 0 0-.512-.8A12.8 12.8 0 0 0 12 5.365a12.8 12.8 0 0 0-5.238 1.068a.877.877 0 0 0-.512.804v0Zm11.5 4.013h5.5"/>`),
		g.Group(children),
	)
}