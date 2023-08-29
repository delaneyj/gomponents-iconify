package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vectorassetcreator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.538 23.739a18.491 18.491 0 0 0-2.264-4.184l2.59-2.551a3.022 3.022 0 0 0-4.282-4.266l-.002.002l-2.55 2.6a18.41 18.41 0 0 0-21.515 0l-2.57-2.57a3.174 3.174 0 0 0-2.171-.916a2.98 2.98 0 0 0-2.122.876a3.018 3.018 0 0 0 0 4.264l2.599 2.55a18.496 18.496 0 0 0-3.476 10.807v1.797m13.63-5.613a2.92 2.92 0 1 1-2.921-2.92a2.92 2.92 0 0 1 2.92 2.92Zm10.547 2.02a2.92 2.92 0 1 1 4.904-1.171"/><circle cx="5.775" cy="34.148" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="42.225" cy="25.614" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.751 33.87l32.531-7.791M7.775 34.148h1.5"/><path fill="none" stroke="currentColor" stroke-dasharray="2.909 2.909" stroke-linecap="round" stroke-linejoin="round" d="M12.184 34.148h27.631"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.269 34.148h1.5v-1.5m-.046-3.578q-.053-.75-.166-1.489"/>`),
		g.Group(children),
	)
}