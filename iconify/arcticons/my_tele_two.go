package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyTeleTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.229 19.697v8.493l3.753.148m-8.932-.352l3.392.134m-3.392-8.214l3.392-.137m-3.392 4.177l2.186-.001m-2.186-4.039v8.08m10.864.429l4.172.164m-4.172-9.11l4.172-.168m-4.172 4.641l2.685-.001m-2.685-4.472v8.946M9.5 20.204l4.051-.159m-2.068 7.878v-7.797m20.757 2.219a3.232 3.232 0 0 1 3.05-3.253a3.113 3.113 0 0 1 3.21 3.166a3.075 3.075 0 0 1-.958 2.325c-1.299 1.143-5.302 4.08-5.302 4.08l6.26.247"/>`),
		g.Group(children),
	)
}