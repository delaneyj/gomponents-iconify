package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stadium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M640 960H64q-26 0-45-18.5T0 896V64q0-26 19-45T64 0h576q26 0 45 19t19 45v832q0 27-18.5 45.5T640 960zm-384-64h192V768H256v128zM448 64H256v128h192V64zm192 0H512v192H192V64H64v384h131q11-55 55-91.5T352 320t102 36.5t55 91.5h131V64zM256 480q0 40 28 68t68 28t68-28t28-68t-28-68t-68-28t-68 28t-28 68zm384 32H509q-11 55-55 91.5T352 640t-102-36.5t-55-91.5H64v384h128V704h320v192h128V512z"/>`),
		g.Group(children),
	)
}