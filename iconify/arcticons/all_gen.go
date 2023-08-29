package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AllGen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.5 37.15l1.95-5.779l1.95 5.779m-.651-1.95h-2.598m9.223-3.829v5.779h2.889m-7.32-5.779v5.779h2.889M28.7 33.321a1.943 1.943 0 0 0-1.936-1.95h-.016h0a1.943 1.943 0 0 0-1.95 1.936v1.965c0 1.083.868 1.878 1.95 1.878h0c1.085 0 1.951-.868 1.951-1.878h-1.95m6.381 1.878h-2.889v-5.779h2.889m-2.889 2.89h1.878m2.552 2.889v-5.779L38.5 37.15v-5.779m-29-5.726h29m-1.179-13.668l1.35-1.895m-4.284 11.042c-5.095.38-4.584-4.759-.486-4.867c2.962-.08 4.591-1.322 4.408-2.826"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.758 13.55c.424-1.282 1.63-2.258 3.059-2.258h1.512c1.785 0 3.222 1.493 3.222 3.278v6.554m-18.838-7.328c.774-1.495 2.242-2.505 3.927-2.505c2.487 0 4.504 2.2 4.504 4.915s-2.017 4.917-4.504 4.917c-2.313 0-4.217-1.901-4.474-4.35h3.404m-6.758-4.481c-.625-.645-1.727-1.097-3-1c-3.773.285-4.751 4.77.467 5.365c2.04.212 3.337 1.047 3.337 2.385c0 1.946-1.923 2.072-4.903 2.072H9.72v-2.397"/>`),
		g.Group(children),
	)
}