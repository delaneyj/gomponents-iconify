package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M704 1536q-40 0-75-15t-61-41t-41-61t-15-75V192q0-40 15-75t41-61t61-41t75-15h512q40 0 75 15t61 41t41 61t15 75v1152q0 40-15 75t-41 61t-61 41t-75 15H704zm-64-192q0 26 19 45t45 19h512q26 0 45-19t19-45V192q0-26-19-45t-45-19H704q-26 0-45 19t-19 45v1152zm1024-320v362q0 84-32 158t-87 129t-129 87t-158 32h-234v128h256v128H640v-128h256v-128H662q-84 0-158-32t-129-87t-87-129t-32-158v-362h128v362q0 57 22 108t59 88t89 60t108 22h596q57 0 108-22t88-59t60-89t22-108v-362h128z"/>`),
		g.Group(children),
	)
}