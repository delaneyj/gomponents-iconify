package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeexpressDriver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.562 23.821a7.8 7.8 0 1 1 13.275-8.195a7.8 7.8 0 0 1-13.275 8.196"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.7 26.479a7.8 7.8 0 1 1 3.9-6.755m-15.599 0H39.6m-31.2-.001V5.5m.843 28.491h2.09m1.115 3.205H9.243v-6.41h3.205m4.813 2.164l-3.205 4.247m3.205 0l-3.205-4.247m4.782 2.644a1.602 1.602 0 1 0 3.205 0v-1.042a1.602 1.602 0 1 0-3.205 0m0-1.602v6.41m5.04-4.808c0-.885.718-1.602 1.603-1.602m-1.603 0v4.246m5.808-.808a1.602 1.602 0 0 1-2.995-.794v-1.041a1.602 1.602 0 1 1 3.205 0v.52h-3.205m4.829 1.765c.292.246.608.359 1.317.359h.36c.585 0 1.059-.476 1.059-1.062s-.474-1.062-1.06-1.062h-.718a1.06 1.06 0 0 1-1.06-1.061h0a1.06 1.06 0 0 1 1.06-1.062h.36c.709 0 1.024.112 1.317.358m1.725 3.53c.292.246.608.359 1.318.359h.36c.584 0 1.058-.476 1.058-1.062s-.474-1.062-1.059-1.062h-.719a1.06 1.06 0 0 1-1.06-1.061h0a1.06 1.06 0 0 1 1.06-1.062h.36c.709 0 1.025.112 1.317.358"/>`),
		g.Group(children),
	)
}