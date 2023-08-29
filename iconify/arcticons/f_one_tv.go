package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FOneTv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.746 27.335h2.79m-1.369 4.001v-4.001m6.01 0l-1.369 4.001l-1.422-4.001" class="a"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.5 16.665h-5.922L27.189 27.96h5.922L44.5 16.665zm-5.922-.001h-20.79a7.03 7.03 0 0 0-4.95 2.033L3.5 27.96h5.922l6.568-6.513a3.515 3.515 0 0 1 2.475-1.017H34.78"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.422 27.96h5.923l3.357-3.33a1.506 1.506 0 0 1 1.06-.435h11.223"/>`),
		g.Group(children),
	)
}