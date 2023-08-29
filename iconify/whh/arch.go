package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 832V704h192v128H832zm0-320v-3l187-64q5 38 5 67v128H832V512zm-46-165l138-138q58 79 83 173l-182 62q-11-51-39-97zM651 224l59-184q98 41 173 120L747 296q-41-45-96-72zm-139-32q-39 0-79 10L374 19Q444 0 512 0t137 19l-58 183q-40-10-79-10zM277 296L141 160q75-79 173-120l59 184q-55 27-96 72zm-78 148L17 382q25-94 83-173l138 138q-28 46-39 97zm-7 68v128H0V512q0-29 5-67l187 64v3zm0 320H0V704h192v128zm0 64v128H0V896h192zm832 0v128H832V896h192z"/>`),
		g.Group(children),
	)
}