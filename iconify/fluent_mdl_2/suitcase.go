package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Suitcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1856 384q40 0 75 15t61 41t41 61t15 75v1152q0 40-15 75t-41 61t-61 41t-75 15H192q-40 0-75-15t-61-41t-41-61t-15-75V576q0-40 15-75t41-61t61-41t75-15h448V256q0-27 10-50t27-40t41-28t50-10h512q27 0 50 10t40 27t28 41t10 50v128h448zm-1088 0h512V256H768v128zM512 512v1280h1024V512H512zM128 1728q0 26 19 45t45 19h192V512H192q-26 0-45 19t-19 45v1152zM1920 576q0-26-19-45t-45-19h-192v1280h192q26 0 45-19t19-45V576z"/>`),
		g.Group(children),
	)
}