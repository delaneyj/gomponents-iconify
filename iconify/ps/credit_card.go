package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 0H64Q37 0 18.5 18.5T0 64v256q0 27 18.5 45.5T64 384h384q27 0 45.5-18.5T512 320V64q0-27-18.5-45.5T448 0zm21 320q0 21-21 21H64q-21 0-21-21v-21h426v21zm0-64H43v-21h426v21zm0-64H43V64q0-21 21-21h384q21 0 21 21v128zM405 85h-85q-21 0-21 22q0 21 21 21h85q10 0 16-6t6-15q0-22-22-22zm-128 22q0 8-6 14.5t-15 6.5t-15-6.5t-6-14.5q0-9 6-15.5t15-6.5t15 6.5t6 15.5z"/>`),
		g.Group(children),
	)
}