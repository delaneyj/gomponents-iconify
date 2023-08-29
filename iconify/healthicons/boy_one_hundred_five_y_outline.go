package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoyOneHundredFiveYOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M34.468 20.948a1 1 0 0 0-1.22-.716c-3.662.954-6.466 1.394-9.244 1.383c-2.782-.012-5.59-.475-9.263-1.385a1 1 0 1 0-.482 1.94c2.162.536 4.068.93 5.873 1.17l.868.116v5.481l-.992 7.939a1 1 0 0 0 1.975.308l1.5-8a1 1 0 0 0 .017-.18h2a1 1 0 0 0 .017.18l1.5 8a1 1 0 0 0 1.975-.308L28 28.938v-5.58l.844-.134c1.518-.24 3.126-.592 4.908-1.056a1 1 0 0 0 .716-1.22Zm-1.724-2.651a3 3 0 0 1 1.512 5.806a56.882 56.882 0 0 1-4.256.954v3.756l.977 7.815a3 3 0 0 1-5.926.925l-.551-2.94l-.551 2.94a3 3 0 0 1-5.926-.925L19 28.813v-3.615c-1.64-.256-3.352-.622-5.222-1.086a3 3 0 0 1 1.444-5.824c3.602.893 6.235 1.316 8.79 1.326c2.546.01 5.167-.389 8.732-1.317Z"/><path d="M24 10a3 3 0 1 0 0 6a3 3 0 0 0 0-6Zm-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z"/></g>`),
		g.Group(children),
	)
}