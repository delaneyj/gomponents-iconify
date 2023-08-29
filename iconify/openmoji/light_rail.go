package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightRail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M68.197 49H18.769c-6.29 0-11.438-4.474-11.438-9.943C7.331 33.59 21.057 20 27.347 20h40.985v28.883a.127.127 0 0 1-.135.117Z"/><path fill="#92d3f5" d="m30.331 38l10.589-9h4.411v9h-15m19 0l10.589-9h4.411v9h-15m-38.999-1s5-7 10-8a22.795 22.795 0 0 1 8 0l-7.57 7.782Z"/><path fill="none" stroke="#ea5a47" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M64.332 43h-52"/><path fill="#3f3f3f" d="M68.331 56h-57l6.235-4h50.765v4z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M68.331 56h-57l6.235-4h50.765v4zm-37-18h14v-8.357m5 8.357h14v-8.357m-36 .357l-7 7h-10"/><path d="M68.197 49H18.769c-6.29 0-11.438-4.474-11.438-9.943h0C7.331 33.59 21.057 20 27.347 20h40.985v28.883a.127.127 0 0 1-.135.117Z"/></g>`),
		g.Group(children),
	)
}