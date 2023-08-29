package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Connectips(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="arcticonsConnectips0" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.682 14.75a9.24 9.24 0 0 0-9.23-9.25h-5.478a9.24 9.24 0 0 0-9.23 9.25h0a9.24 9.24 0 0 0 9.23 9.25h6.052a9.24 9.24 0 0 1 9.23 9.25h0a9.24 9.24 0 0 1-9.23 9.25h-5.478"/><path id="arcticonsConnectips1" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.25 35.682a9.24 9.24 0 0 0 9.25-9.23v-5.478a9.24 9.24 0 0 0-9.25-9.23h0a9.24 9.24 0 0 0-9.25 9.23v6.052a9.24 9.24 0 0 1-9.25 9.23h0a9.24 9.24 0 0 1-9.25-9.23v-5.478"/></defs><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.925 15.998l2.085 3.288l3.288-2.084"/><use href="#arcticonsConnectips0" stroke-linecap="round" stroke-linejoin="round"/><use href="#arcticonsConnectips0" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.682 14.75a13.62 13.62 0 0 1-.672 4.536M15.075 32.002l-2.085-3.288l-3.288 2.084m2.616 2.452a9.24 9.24 0 0 0 9.23 9.25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.318 33.25a9.24 9.24 0 0 0 9.23 9.25m-9.23-9.25a13.62 13.62 0 0 1 .672-4.536m19.012 4.211l-3.288 2.085l2.084 3.288"/><use href="#arcticonsConnectips1" stroke-linecap="round" stroke-linejoin="round"/><use href="#arcticonsConnectips1" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.25 35.682a13.62 13.62 0 0 1-4.536-.672M15.998 15.075l3.288-2.085l-2.084-3.288m-2.452 2.616a9.24 9.24 0 0 0-9.25 9.23"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.75 12.318a9.24 9.24 0 0 0-9.25 9.23m9.25-9.23a13.62 13.62 0 0 1 4.536.672"/>`),
		g.Group(children),
	)
}