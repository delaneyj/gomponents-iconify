package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Termius(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.295 35.024A9.803 9.803 0 0 1 4.5 25.212a10.187 10.187 0 0 1 4.535-8.397a9.867 9.867 0 0 1 9.82-9.945a10.165 10.165 0 0 1 7.406 3.272a10.524 10.524 0 0 1 4.336-.963a9.697 9.697 0 0 1 9.733 9.748a9.751 9.751 0 0 1-7.032 16.936a11.07 11.07 0 0 1-9.14 5.267a10.015 10.015 0 0 1-9.141-6.073a4.954 4.954 0 0 0-.721-.033Zm2.98-14.3l5.87 1.86m-5.87 1.86l5.87-1.86"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.919 26.975a5.9 5.9 0 0 0 6.302 0"/>`),
		g.Group(children),
	)
}