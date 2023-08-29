package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PcloudPass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M36.05 24.174c.192-6.69-5.076-12.27-11.767-12.463s-12.27 5.076-12.463 11.767a12.06 12.06 0 0 0 0 .696"/><path d="M13.242 18.115a9.1 9.1 0 0 0 .358 18.18h24a5.84 5.84 0 0 0 3.48-10.53h.05a5.999 5.999 0 0 0 .24-1.59a5.85 5.85 0 0 0-5.85-5.85a6.712 6.712 0 0 0-.75.063"/></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="19.508" cy="24.674" r="3.508"/><path d="M32 28.182v-3.508h0h-8.989m6.687 2.023v-2.023"/></g>`),
		g.Group(children),
	)
}