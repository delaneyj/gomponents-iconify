package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorkCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 91h-85V69q0-27-18.5-45.5T299 5h-86q-27 0-45.5 18.5T149 69v22H64q-27 0-45.5 18T0 155v85q0 21 21 21v171q0 17 13 30t30 13h384q17 0 30-13t13-30V261q21 0 21-21v-85q0-28-18.5-46T448 91zM192 69q0-9 6-15t15-6h86q9 0 15 6t6 15v22H192V69zM64 432V261h128v64q0 18 12.5 30.5T235 368h42q18 0 30.5-12.5T320 325v-64h128v171H64zm171-107v-64h42v64h-42zm234-106H43v-64q0-22 21-22h384q21 0 21 22v64z"/>`),
		g.Group(children),
	)
}