package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spades(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12.775 5.44C9.751 3.192 8.708 1.393 8.001 0c-.708 1.393-1.75 3.192-4.774 5.44c-5.157 3.833-.303 9.182 3.965 6.238c-.278 1.827-1.227 3.159-2.191 3.733v.59h6v-.59c-.964-.574-1.913-1.906-2.191-3.733c4.268 2.944 9.122-2.405 3.965-6.238z"/>`),
		g.Group(children),
	)
}