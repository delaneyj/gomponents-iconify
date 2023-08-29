package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Juegos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.753 26.186V15.111m-5.537 5.537h11.075m14.831 3.598l-5.926-3.421v6.843l5.926-3.422zm-4.939-6.089l5.926 3.422v-6.843l-5.926 3.421z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 33.18h6.28s4.954 5.44 8.792 5.44s4.517-3.06 4.42-5.756s-.87-12.37-1.19-16.49c-.21-2.698-2.429-6.995-8.258-6.995a8.582 8.582 0 0 0-5.948 2.361h-8.192a8.582 8.582 0 0 0-5.948-2.36c-5.829 0-8.047 4.296-8.257 6.994c-.321 4.12-1.093 13.795-1.19 16.49s.582 5.756 4.42 5.756s8.791-5.44 8.791-5.44Z"/>`),
		g.Group(children),
	)
}