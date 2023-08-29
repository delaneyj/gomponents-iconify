package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionPdf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H128q-18 0-30.5-12.5T85 301V45q0-17 12.5-29.5T128 3h256zM203 163v-22q0-13-9.5-22.5T171 109h-54v128h32v-42h22q13 0 22.5-9.5T203 163zm106 42v-64q0-13-9-22.5t-23-9.5h-53v128h53q14 0 23-9t9-23zm86-64v-32h-64v128h32v-42h32v-32h-32v-22h32zm-246 22v-22h22v22h-22zM43 88v299h298v42H43q-18 0-30.5-12.5T0 387V88h43zm213 117v-64h21v64h-21z"/>`),
		g.Group(children),
	)
}