package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plurk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.501 16.641a8.887 8.887 0 0 1-.235-1.135a2.001 2.001 0 1 1 3.761.953l2.8 2.104a1.7 1.7 0 0 1 .567 1.59c-.11.89-.736 1.549-1.398 1.467a.937.937 0 0 1-.441-.177l-3.111-2.318A2.001 2.001 0 1 1 5.5 16.64Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.504 31.434A16.903 16.903 0 0 1 8.793 24c0-.247.005-.494.015-.738a3.22 3.22 0 0 0 2.414 2.384c1.884.23 3.712-2.026 4.082-5.038s-.859-5.644-2.741-5.875a2.296 2.296 0 0 0-1.112.146a16.98 16.98 0 0 1 30.161 2.999c.966-.022 2.587-1.193 2.787-.881c.252.39.174 2.572-1.927 3.923a17.073 17.073 0 0 1 .28 3.08a16.889 16.889 0 0 1-2.33 8.583l1.33 1.328a.956.956 0 0 1 .002 1.35l-3.46 3.469a.956.956 0 0 1-1.35.001l-1.08-1.078a16.967 16.967 0 0 1-21.22-.832l-1.354 1.137a.956.956 0 0 1-1.346-.118l-3.149-3.752a.955.955 0 0 1 .118-1.345s1.533-1.278 1.591-1.309Z"/>`),
		g.Group(children),
	)
}