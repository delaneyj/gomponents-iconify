package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DsdSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.144 5.785L9.866 10.45a5.275 5.275 0 0 0-3.426 5.044v17.132a5.359 5.359 0 0 0 3.521 5.045l12.183 4.473a5.003 5.003 0 0 0 3.712 0l12.182-4.473a5.359 5.359 0 0 0 3.522-5.045V15.493a5.508 5.508 0 0 0-3.426-5.044L25.95 5.786a6.488 6.488 0 0 0-3.807 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.56 15.493a2.564 2.564 0 0 0-3.522-2.474l-12.182 4.378a5.003 5.003 0 0 1-3.712 0L10.056 13.02a2.704 2.704 0 0 0-3.616 2.57m5.037 3.138l6.555 2.341a4.6 4.6 0 0 1 3.053 4.332v8.92a2 2 0 0 1-2.673 1.884l-6.94-2.479a1 1 0 0 1-.663-.942V19.2a.5.5 0 0 1 .668-.471Zm26.04-.448l-9.613 3.434a1 1 0 0 0-.663.942v13.586a.5.5 0 0 0 .668.47l8.945-3.194a1 1 0 0 0 .664-.942V26.49a.5.5 0 0 0-.669-.47l-9.608 3.43"/>`),
		g.Group(children),
	)
}