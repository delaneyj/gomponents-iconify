package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tiddloid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.595 12.44c-.392 2.743.224 6.941 5.766 6.941c2.855 0 3.688-.591 3.688-.591"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.647 24.878A25.61 25.61 0 0 0 25 27.107c2.911 0 5.43.112 8.033.112a8.036 8.036 0 0 0 2.687-1.175c1.288-.896 2.463-3.639 2.463-3.639c.98.54 2.141.678 3.618-.837c1.608-1.647.916-5.264.916-5.936a14.312 14.312 0 0 1 .784-2.743a7.834 7.834 0 0 1-4.198 1.567c-2.103-.17-3.471-2.35-4.143-2.35s-1.527 2.815-1.303 5.502a67.103 67.103 0 0 0-7.821-.409c-1.52.075-3.75-1.287-6.438-1.287s-6.101 1.343-6.101 5.374a3.816 3.816 0 0 0 2.151 3.594a8.116 8.116 0 0 0-1.256 5.082c0 1.652 2.435 4.283 3.051 4.283h1.064"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.528 26.26a8.719 8.719 0 0 0-.938 4.066c0 1.231 1.903 3.526 3.134 4.058c1.04.205 1.866.299 1.866.299m8.127-7.473a46.415 46.415 0 0 0-.047 4.768a15.182 15.182 0 0 0 2.957 3.665a2.273 2.273 0 0 0 1.064.252m-5.907-8.719a15.213 15.213 0 0 0-.95 4c.083.504 1.057 3.506 1.59 3.786a2.475 2.475 0 0 0 1.23.205"/>`),
		g.Group(children),
	)
}