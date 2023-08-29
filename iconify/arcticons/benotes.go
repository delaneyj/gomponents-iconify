package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Benotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 24C2.5 12.1 12.1 2.5 24 2.5S45.5 12.1 45.5 24S35.9 45.5 24 45.5S2.5 35.9 2.5 24Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 22.443h5.75v-6.828H12.5v6.828Zm0 9.942h5.75v-6.828H12.5v6.828Zm8.625 0h5.75v-6.828h-5.75v6.828Zm8.625 0h5.75v-6.828h-5.75v6.828Zm-8.625-9.942h5.75v-6.828h-5.75v6.828Zm8.625-6.828v6.828h5.75v-6.828h-5.75Z"/>`),
		g.Group(children),
	)
}