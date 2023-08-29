package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jupiter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.448 20.19v6.673a2.224 2.224 0 0 1-2.224 2.224h0A2.224 2.224 0 0 1 9 26.863v-.739m6.406-4.337v3.011c0 1.008.817 1.825 1.825 1.825h0a1.825 1.825 0 0 0 1.825-1.825v-3.01m0 3.01v1.825m18.119-3.011c0-1.008.817-1.825 1.825-1.825h0m-1.825 0v4.836"/><circle cx="26.663" cy="19.552" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.663 21.787v4.836m2.794-6.341v6.341m-.958-4.836h1.916m4.7 3.915c-.314.55-.907.921-1.586.921h0a1.825 1.825 0 0 1-1.825-1.825v-1.186c0-1.008.817-1.825 1.825-1.825h0c1.008 0 1.825.817 1.825 1.825v.593h-3.65m-10.67.593c0 1.008.817 1.825 1.825 1.825h0a1.825 1.825 0 0 0 1.825-1.825v-1.186a1.825 1.825 0 0 0-1.825-1.825h0a1.825 1.825 0 0 0-1.825 1.825m0-1.825v7.3"/>`),
		g.Group(children),
	)
}