package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Miui(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 640V0h192v640H832zM512 256q0-27-18.5-45.5T448 192H192v448H0V0h512q80 0 136 56t56 136v448H512V256zm-64 384H256V256h192v384z"/>`),
		g.Group(children),
	)
}