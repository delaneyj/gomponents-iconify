package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoPro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.866 14.701v-3.196h-7.58V14.7"/><circle cx="26.956" cy="25.924" r="12.7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.663 37.876H6.775a2.26 2.26 0 0 1-2.26-2.26h0V13.971h18.148m8.586 0h12.265v21.645a2.26 2.26 0 0 1-2.26 2.26H31.25m11.465-23.905v-3.493h-8.284v3.493m-.001 0v-3.16L31.943 6.5H21.97l-2.489 4.31v3.161"/><circle cx="26.956" cy="25.927" r="6.572" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="26.956" cy="25.927" r="10.173" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.43 11.417H19.482"/>`),
		g.Group(children),
	)
}