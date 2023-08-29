package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Palmstore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.084 38.17l17.302 4.383a3.244 3.244 0 0 0 4.04-3.145V8.592a3.244 3.244 0 0 0-4.04-3.145L8.084 9.83A3.424 3.424 0 0 0 5.5 13.15v21.7a3.424 3.424 0 0 0 2.584 3.32Zm21.343.874l3.1.786a2.768 2.768 0 0 0 3.448-2.684V10.854a2.768 2.768 0 0 0-3.448-2.683l-3.1.785"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.975 36.037l3.7.937a2.269 2.269 0 0 0 2.825-2.199v-21.55a2.269 2.269 0 0 0-2.826-2.2l-3.7.938M10.41 38.759V17.203l6.28-.665a5.944 5.944 0 0 1 6.568 5.91h0c0 3.372-2.8 6.068-6.169 5.94l-6.68-.253"/>`),
		g.Group(children),
	)
}