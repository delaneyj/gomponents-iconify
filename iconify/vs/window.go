package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Window(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M64 0Q38 0 19 19T0 64v1536q0 26 19 45t45 19h1536q26 0 45-19t19-45V64q0-26-19-45t-45-19H64zm1408 896v576H896V896h576zm-704 576H192V896h576v576zM896 192h576v576H896V192zm-704 0h576v576H192V192zm960 384q33 55 124 91.5t132 36.5V256h-384q27 116 51 175t77 145zm-640 0q53-86 77-145t51-175H256v448q41 0 132-36.5T512 576zm832 448q-26 66-44.5 171t-19.5 213h128V960q-24 0-37 13.5t-27 50.5zm-1024 0q-14-37-27-50.5T256 960v448h128q-1-108-19.5-213T320 1024z"/>`),
		g.Group(children),
	)
}