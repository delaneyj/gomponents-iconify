package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenTime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 1024q63 83 95 181t33 203q0 133-50 249t-137 204t-203 137t-250 50q-133 0-249-50t-204-137t-137-203t-50-250H128q-27 0-50-10t-40-27t-28-41t-10-50V256q0-27 10-50t27-40t41-28t50-10h1664q27 0 50 10t40 27t28 41t10 50v768zM781 1280q22-112 80-206t142-162t187-106t218-38q104 0 202 32t182 96V256H128v1024h653zm627 640q106 0 199-40t162-110t110-163t41-199q0-106-40-199t-110-162t-163-110t-199-41q-106 0-199 40t-162 110t-110 163t-41 199q0 106 40 199t110 162t163 110t199 41zm0-512h256v128h-384v-512h128v384z"/>`),
		g.Group(children),
	)
}