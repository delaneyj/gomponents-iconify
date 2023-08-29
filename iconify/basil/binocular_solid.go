package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BinocularSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.096 4.91c.354.302.665.66.904 1.061c.239-.401.55-.759.904-1.061c.816-.696 1.901-1.132 2.985-1.132c.758 0 1.57.214 2.27.576c.698.361 1.334.895 1.682 1.572a.525.525 0 0 1 .023.051l2.13 5.6a5.194 5.194 0 1 1-9.609 3.2h-.77a5.195 5.195 0 1 1-9.609-3.2l2.13-5.6a.504.504 0 0 1 .023-.05c.349-.678.984-1.212 1.681-1.573c.701-.362 1.513-.576 2.271-.576c1.084 0 2.169.436 2.985 1.132ZM2.75 14.278a3.694 3.694 0 1 1 7.389 0a3.694 3.694 0 0 1-7.389 0Zm11.111 0a3.694 3.694 0 1 1 7.389 0a3.694 3.694 0 0 1-7.389 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}