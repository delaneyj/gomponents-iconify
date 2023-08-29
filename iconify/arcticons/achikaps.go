package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Achikaps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.49 5.5H7.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2.01-2Z"/><circle cx="11.88" cy="12.18" r="2.48" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="21.31" cy="12.72" r="1.68" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.82" cy="10.54" r="1.68" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="37.66" cy="31.9" r="1.68" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.74" cy="37.3" r="1.68" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="22" cy="38.55" r="1.68" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="13.25" cy="33.56" r="1.68" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="9.15" cy="23.56" r="1.68" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.43 9l1.32-3.46m2.39 13.11L32.68 12m5.73 7.59l4.09-1.82m-3.32 14.86l3.31 1.68m-6.37-3.08l-3.44-1.62m3.61 3.27l-4.18 3.41m-.59 2.5l2.11 3.71m-9.97-4.22l5.43-.65m-7.32 2.6v2.27m-9.54-7.59l-5.57 7.41M22.94 12.3l7.23-1.39m-10.53 1.67l-5.28-.15M11 9.84L9.39 5.5m28.08 24.72L37 21.86m-16.44-7.65l-2 3.78m-7.77 5.17l2.61-.66m-5.9.73l-2-.49m8.22 9.2L14.2 30m6.31 7.77l-5.82-3.34m11.39-16.85L22.5 13.9m4.46 5.9v8.12m.04-3.35a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2v3.35"/><circle cx="33.09" cy="20.06" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.09 22.54v5.38m-8.24-1.02a2 2 0 0 1-1.76 1h0a2 2 0 0 1-2-2v-1.33a2 2 0 0 1 2-2h0a2 2 0 0 1 1.76 1M14.2 27.89l2.69-8.09m2.59 8.12l-2.59-8.12m1.72 5.4H15.1"/><circle cx="36.81" cy="20.2" r="1.68" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.84 14.43l-1.69 7.44m0 3.37l2.95 7.1"/>`),
		g.Group(children),
	)
}