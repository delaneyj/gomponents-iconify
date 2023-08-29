package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ringer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 1536v128h-512q0 53-20 99t-55 82t-81 55t-100 20q-53 0-99-20t-82-55t-55-81t-20-100H256v-128h128V768q0-88 23-170t64-153t100-129t130-100t153-65t170-23q88 0 170 23t153 64t129 100t100 130t65 153t23 170v768h128zm-256 0V768q0-106-40-199t-110-162t-163-110t-199-41q-106 0-199 40T663 406T553 569t-41 199v768h1024zm-512 256q27 0 50-10t40-27t28-41t10-50H896q0 27 10 50t27 40t41 28t50 10z"/>`),
		g.Group(children),
	)
}