package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Funmax(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.25 22.611v-5.427m-2.713 2.713h5.426m12.081 6.754l4.553 6.83a5.405 5.405 0 0 0 9.786-4.107c-.984-4.67-1.769-9.026-2.572-11.982c-.846-3.113-3.884-5.28-7.284-5.28c-2.052 0-3.912.81-5.288 2.123h-8.478a7.635 7.635 0 0 0-5.287-2.124c-3.401 0-6.44 2.168-7.285 5.28c-.803 2.957-1.588 7.313-2.572 11.983a5.405 5.405 0 0 0 9.787 4.107l4.553-6.83h10.087Z"/><circle cx="31.03" cy="19.897" r=".75" fill="currentColor"/><circle cx="36.463" cy="19.897" r=".75" fill="currentColor"/><circle cx="33.75" cy="17.184" r=".75" fill="currentColor"/><circle cx="33.75" cy="22.611" r=".785" fill="currentColor"/>`),
		g.Group(children),
	)
}