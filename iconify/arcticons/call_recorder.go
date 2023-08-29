package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallRecorder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.03 32.59v4.752M45.5 24c0 11.307-9.166 20.473-20.473 20.473a20.38 20.38 0 0 1-9.901-2.55L2.5 44.09l4.06-11.245A20.392 20.392 0 0 1 4.554 24c0-11.307 9.166-20.473 20.473-20.473S45.5 12.694 45.5 24Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.095 9.274c2.136-.02 4.914 1.942 5.071 4.781V24.85c0 2.8-2.27 5.071-5.071 5.071h0a5.07 5.07 0 0 1-5.071-5.071V14.055c0-2.8 2.935-4.763 5.07-4.781Zm-5.072 5.693l2.288.027m-2.288 2.579l2.288.027m-2.288 2.578l2.288.027"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.073 20.721l.124 3.748c.171 5.187 4.082 8.325 7.898 8.325s7.89-3.604 8.007-7.742l.12-4.304m-1.14 16.909H17.86"/>`),
		g.Group(children),
	)
}