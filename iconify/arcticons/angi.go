package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.464 39.567l5.895.053L32.61 4.5h-8.584l-2.663 18.146c-.214 1.97-1.725 15.785-4.31 15.238c-5.07-3.454.625-9.686 3.088-9.28c9.285-1.3 11.791 2.027 12.324 10.963Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.465 9.954l-1.94 13.23a8.146 8.146 0 0 1 2.602.726a13.89 13.89 0 0 1 1.363.76ZM14.82 42.677a10.48 10.48 0 0 1-4.631-12.436a9.709 9.709 0 0 1 10.997-6.696"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.682 42.606c8.203 4.104 10.078-6.857 11.131-14.002"/>`),
		g.Group(children),
	)
}